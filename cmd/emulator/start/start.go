/*
 * Flow Emulator
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package start

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/onflow/cadence"
	sdk "github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/prometheus/common/log"
	"github.com/psiemens/sconfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/onflow/flow-emulator/server"
)

type Config struct {
	Port                   int           `default:"3569" flag:"port,p" info:"port to run RPC server"`
	HTTPPort               int           `default:"8080" flag:"http-port" info:"port to run HTTP server"`
	Verbose                bool          `default:"false" flag:"verbose,v" info:"enable verbose logging"`
	BlockTime              time.Duration `flag:"block-time,b" info:"time between sealed blocks, e.g. '300ms', '-1.5h' or '2h45m'. Valid units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'"`
	ServicePrivateKey      string        `flag:"service-priv-key" info:"service account private key"`
	ServicePublicKey       string        `flag:"service-pub-key" info:"service account public key"`
	ServiceKeySigAlgo      string        `flag:"service-sig-algo" info:"service account key signature algorithm"`
	ServiceKeyHashAlgo     string        `flag:"service-hash-algo" info:"service account key hash algorithm"`
	Init                   bool          `default:"false" flag:"init" info:"whether to initialize a new account profile"`
	GRPCDebug              bool          `default:"false" flag:"grpc-debug" info:"enable gRPC server reflection for debugging with grpc_cli"`
	Persist                bool          `default:"false" flag:"persist" info:"enable persistent storage"`
	DBPath                 string        `default:"./flowdb" flag:"dbpath" info:"path to database directory"`
	SimpleAddresses        bool          `default:"false" flag:"simple-addresses" info:"use sequential addresses starting with 0x01"`
	TokenSupply            string        `default:"100000000000.0" flag:"token-supply" info:"initial FLOW token supply"`
	TransactionExpiry      int           `default:"10" flag:"transaction-expiry" info:"transaction expiry, measured in blocks"`
	StorageLimitEnabled    bool          `default:"true" flag:"storage-limit" info:"enable account storage limit"`
	TransactionFeesEnabled bool          `default:"false" flag:"transaction-fees" info:"enable transaction fees"`
	TransactionMaxGasLimit int           `default:"9999" flag:"transaction-max-gas-limit" info:"maximum gas limit for transactions"`
	ScriptGasLimit         int           `default:"100000" flag:"script-gas-limit" info:"gas limit for scripts"`
}

const EnvPrefix = "FLOW"

var (
	logger *logrus.Logger
	conf   Config
)

type serviceKeyFunc func(
	init bool,
	sigAlgo crypto.SignatureAlgorithm,
	hashAlgo crypto.HashAlgorithm,
) (crypto.PrivateKey, crypto.SignatureAlgorithm, crypto.HashAlgorithm)

func Cmd(getServiceKey serviceKeyFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts the Flow emulator server",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				servicePrivateKey  crypto.PrivateKey
				servicePublicKey   crypto.PublicKey
				serviceKeySigAlgo  crypto.SignatureAlgorithm
				serviceKeyHashAlgo crypto.HashAlgorithm
				err                error
			)

			serviceKeySigAlgo = crypto.StringToSignatureAlgorithm(conf.ServiceKeySigAlgo)
			serviceKeyHashAlgo = crypto.StringToHashAlgorithm(conf.ServiceKeyHashAlgo)

			if len(conf.ServicePublicKey) > 0 {
				checkKeyAlgorithms(serviceKeySigAlgo, serviceKeyHashAlgo)

				servicePublicKey, err = crypto.DecodePublicKeyHex(serviceKeySigAlgo, conf.ServicePublicKey)
				if err != nil {
					Exit(1, err.Error())
				}
			} else if len(conf.ServicePrivateKey) > 0 {
				checkKeyAlgorithms(serviceKeySigAlgo, serviceKeyHashAlgo)

				servicePrivateKey, err = crypto.DecodePrivateKeyHex(serviceKeySigAlgo, conf.ServicePrivateKey)
				if err != nil {
					Exit(1, err.Error())
				}

				servicePublicKey = servicePrivateKey.PublicKey()
			} else {
				servicePrivateKey, serviceKeySigAlgo, serviceKeyHashAlgo = getServiceKey(
					conf.Init,
					serviceKeySigAlgo,
					serviceKeyHashAlgo,
				)
				servicePublicKey = servicePrivateKey.PublicKey()
			}

			if conf.Verbose {
				logger.SetLevel(logrus.DebugLevel)
			}

			serviceAddress := sdk.ServiceAddress(sdk.Emulator)
			serviceFields := logrus.Fields{
				"serviceAddress":  serviceAddress.Hex(),
				"servicePubKey":   hex.EncodeToString(servicePublicKey.Encode()),
				"serviceSigAlgo":  serviceKeySigAlgo,
				"serviceHashAlgo": serviceKeyHashAlgo,
			}

			if servicePrivateKey != (crypto.PrivateKey{}) {
				serviceFields["servicePrivKey"] = hex.EncodeToString(servicePrivateKey.Encode())
			}

			logger.WithFields(serviceFields).Infof("⚙️   Using service account 0x%s", serviceAddress.Hex())

			serverConf := &server.Config{
				GRPCPort:  conf.Port,
				GRPCDebug: conf.GRPCDebug,
				HTTPPort:  conf.HTTPPort,
				// TODO: allow headers to be parsed from environment
				HTTPHeaders:            nil,
				BlockTime:              conf.BlockTime,
				ServicePublicKey:       servicePublicKey,
				ServiceKeySigAlgo:      serviceKeySigAlgo,
				ServiceKeyHashAlgo:     serviceKeyHashAlgo,
				Persist:                conf.Persist,
				DBPath:                 conf.DBPath,
				GenesisTokenSupply:     parseTokenSupply(conf.TokenSupply),
				TransactionMaxGasLimit: uint64(conf.TransactionMaxGasLimit),
				ScriptGasLimit:         uint64(conf.ScriptGasLimit),
				TransactionExpiry:      uint(conf.TransactionExpiry),
				StorageLimitEnabled:    conf.StorageLimitEnabled,
				TransactionFeesEnabled: conf.TransactionFeesEnabled,
			}

			emu := server.NewEmulatorServer(logger, serverConf)
			emu.Start()
		},
	}

	initConfig(cmd)

	return cmd
}

func init() {
	initLogger()
}

func initLogger() {
	logger = logrus.New()
	logger.Formatter = new(logrus.TextFormatter)
	logger.Out = os.Stdout
}

func initConfig(cmd *cobra.Command) {
	err := sconfig.New(&conf).
		FromEnvironment(EnvPrefix).
		BindFlags(cmd.PersistentFlags()).
		Parse()
	if err != nil {
		log.Fatal(err)
	}
}

func Exit(code int, msg string) {
	fmt.Println(msg)
	os.Exit(code)
}

func parseTokenSupply(supply string) cadence.UFix64 {
	tokenSupply, err := cadence.NewUFix64(supply)
	if err != nil {
		Exit(
			1,
			fmt.Sprintf(
				"Invalid token supply. Failed to parse `%s` as an unsigned 64-bit fixed-point number: %s",
				conf.TokenSupply,
				err.Error()),
		)
	}

	return tokenSupply
}

func checkKeyAlgorithms(sigAlgo crypto.SignatureAlgorithm, hashAlgo crypto.HashAlgorithm) {
	if sigAlgo == crypto.UnknownSignatureAlgorithm {
		Exit(1, "Must specify service key signature algorithm (e.g. --service-sig-algo=ECDSA_P256)")
	}

	if hashAlgo == crypto.UnknownHashAlgorithm {
		Exit(1, "Must specify service key hash algorithm (e.g. --service-hash-algo=SHA3_256)")
	}
}
