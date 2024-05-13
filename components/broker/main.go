// Copyright 2023 XigXog
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"flag"
	"runtime"
	"time"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/xigxog/kubefox/api"
	"github.com/xigxog/kubefox/components/broker/config"
	"github.com/xigxog/kubefox/components/broker/engine"
	"github.com/xigxog/kubefox/logkf"
	"github.com/xigxog/kubefox/utils"
)

func main() {
	flag.StringVar(&config.Instance, "instance", "", "KubeFox instance Broker is part of. (required)")
	flag.StringVar(&config.Platform, "platform", "", "Platform instance Broker if part of. (required)")
	flag.StringVar(&config.Namespace, "namespace", "", "Namespace of Platform instance. (required)")
	flag.StringVar(&config.GRPCSrvAddr, "grpc-addr", "127.0.0.1:6060", "Address and port the gRPC server should bind to.")
	flag.StringVar(&config.HealthSrvAddr, "health-addr", "127.0.0.1:1111", `Address and port the HTTP health server should bind to, set to "false" to disable.`)
	flag.StringVar(&config.NATSAddr, "nats-addr", "127.0.0.1:4222", "Address and port of NATS server.")
	flag.StringVar(&config.VaultURL, "vault-url", "https://127.0.0.1:8200", "URL of Vault server.")
	flag.StringVar(&config.TelemetryAddr, "telemetry-addr", "127.0.0.1:4318", `Address and port of OTEL telemetry collector, set to "false" to disable.`)
	flag.DurationVar(&config.TelemetryInterval, "telemetry-interval", time.Minute, `Interval at which to report metrics, , set to "0" to disable.`)
	flag.Int64Var(&config.MaxEventSize, "max-event-size", api.DefaultMaxEventSizeBytes, "Maximum size of event in bytes.")
	flag.IntVar(&config.NumWorkers, "num-workers", runtime.NumCPU(), "Number of worker threads to start, default is number of logical CPUs.")
	flag.StringVar(&config.LogFormat, "log-format", "console", `Log format; one of ["json", "console"].`)
	flag.StringVar(&config.LogLevel, "log-level", "debug", `Log level; one of ["debug", "info", "warn", "error"].`)
	flag.Parse()

	utils.CheckRequiredFlag("instance", config.Instance)
	utils.CheckRequiredFlag("platform", config.Platform)
	utils.CheckRequiredFlag("namespace", config.Namespace)

	defer logkf.Global.Sync()
	logkf.Global = logkf.
		BuildLoggerOrDie(config.LogFormat, config.LogLevel).
		WithInstance(config.Instance).
		WithPlatform(config.Platform).
		WithPlatformComponent(api.PlatformComponentBroker)

	engine.New().Start()
}
