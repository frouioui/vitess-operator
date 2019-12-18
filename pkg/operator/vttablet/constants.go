/*
Copyright 2019 PlanetScale Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vttablet

import (
	"time"
)

const (
	vttabletPriorityClassName = "vitess"

	vttabletContainerName = "vttablet"
	vttabletCommand       = "/vt/bin/vttablet"

	vtbackupContainerName = "vtbackup"
	vtbackupCommand       = "/vt/bin/vtbackup"

	mysqldContainerName = "mysqld"
	mysqldCommand       = "/vt/bin/mysqlctld"

	mysqldExporterContainerName      = "mysqld-exporter"
	mysqldExporterCommand            = "/bin/mysqld_exporter"
	mysqldExporterUser               = "vt_dba"
	mysqldExporterPort               = 9104
	mysqldExporterPortName           = "metrics"
	mysqldExporterCPURequestMillis   = 10
	mysqldExporterCPULimitMillis     = 100
	mysqldExporterMemoryRequestBytes = 32 * (1 << 20)  // 32 MiB
	mysqldExporterMemoryLimitBytes   = 128 * (1 << 20) // 128 MiB

	serviceMap          = "grpc-queryservice,grpc-tabletmanager,grpc-updatestream"
	runAsUser           = 999
	fsGroup             = 999
	healthCheckInterval = 5 * time.Second

	grpcMaxMessageSize = 64 * 1024 * 1024 // 64 MiB

	queryserverConfigMaxResultSize = 100000
	queryserverConfigQueryTimeout  = 900

	queryserverConfigPoolSize       = 96
	queryserverConfigStreamPoolSize = 96
	queryserverConfigTransactionCap = 300

	vtRootPath         = "/vt"
	vtBinPath          = vtRootPath + "/bin"
	vtConfigPath       = vtRootPath + "/config"
	vtMycnfPath        = vtConfigPath + "/mycnf"
	vtDataRootPath     = vtRootPath + "/vtdataroot"
	vtMysqlRootPath    = "/usr"
	vtRootVolumeName   = "vt-root"
	mysqlctlSocketPath = vtDataRootPath + "/mysqlctl.sock"
	sslCertsPath       = "/etc/ssl/certs"

	pvcVolumeName = "persistent-volume-claim"

	dbConfigCharset       = "utf8"
	dbConfigAppUname      = "vt_app"
	dbConfigDbaUname      = "vt_dba"
	dbConfigReplUname     = "vt_repl"
	dbConfigFilteredUname = "vt_filtered"

	vreplicationTabletType = "master"

	shardConfigPath      = "/shard-config"
	dbInitScriptFileName = "init_db.sql"
	dbInitScriptFilePath = shardConfigPath + "/" + dbInitScriptFileName

	externalDatastoreCredentialsPath       = "/external-datastore/credentials"
	externalDatastoreCredentialsFilename   = "credentials.json"
	externalDatastoreCredentialsFilePath   = externalDatastoreCredentialsPath + "/" + externalDatastoreCredentialsFilename
	externalDatastoreCredentialsVolumeName = "external-credentials"

	externalDatastoreSSLCAPath       = "/external-certs/ssl-certs"
	externalDatastoreSSLCAFilename   = "ssl-ca.pem"
	externalDatastoreSSLCAFilePath   = externalDatastoreSSLCAPath + "/" + externalDatastoreSSLCAFilename
	externalDatastoreSSLCAVolumeName = "external-datastore-certs"
	enableSSLBitflag                 = 2048

	mysqldConfigOverridesAnnotationName      = "planetscale.com/mysqld-config-overrides"
	mysqldConfigOverridesAnnotationFieldPath = "metadata.annotations['" + mysqldConfigOverridesAnnotationName + "']"

	vtbackupTimeout            = 2 * time.Hour
	vtbackupReplicationTimeout = 1 * time.Hour
	// waitForBackupInterval is how often to poll for new backups when a tablet
	// starts up and finds that no backup exists yet.
	waitForBackupInterval = 10 * time.Second

	// restoreConcurrency is how many files to restore concurrently in vttablet.
	restoreConcurrency = 10
	// vtbackupConcurrency is how many files to backup/restore concurrently in vtbackup.
	vtbackupConcurrency = 10

	homeDir = "/home/vitess"

	xtrabackupEngineName  = "xtrabackup"
	xtrabackupStreamMode  = "xbstream"
	xtrabackupStripeCount = 8
	xtrabackupUser        = "vt_dba"
)

var (
	defaultExtraMyCnf = []string{
		vtMycnfPath + "/rbr.cnf",
	}
)