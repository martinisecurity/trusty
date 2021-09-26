package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-phorce/dolly/netutil"
	"github.com/go-phorce/dolly/xlog"
	"github.com/juju/errors"
	"github.com/martinisecurity/trusty/pkg/configloader"
)

var logger = xlog.NewPackageLogger("github.com/martinisecurity/trusty", "config")

const (
	// ConfigFileName is default name for the configuration file
	ConfigFileName = "trusty-config.yaml"
)

// DefaultFactory returns default configuration factory
func DefaultFactory() (*configloader.Factory, error) {
	var err error

	nodeInfo, err := netutil.NewNodeInfo(nil)
	if err != nil {
		return nil, errors.Trace(err)
	}

	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// try the list of allowed locations to find the config file
	searchDirs := []string{
		cwd,
		filepath.Dir(cwd) + "/etc/dev", // $PWD/etc/dev for running locally on dev machine
		"/opt/trusty/etc/prod",
		"/opt/trusty/etc/stage",
		"/opt/trusty/etc/dev", // for CI test or stage the etc/dev must be after etc/prod
	}

	logger.Infof("searchDirs=[%s]", strings.Join(searchDirs, ","))

	return configloader.NewFactory(nodeInfo, searchDirs, "TRUSTY_")
}

// Load will load the configuration from the named config file,
// apply any overrides, and resolve relative directory locations.
func Load(configFile string) (*Configuration, error) {
	return LoadForHostName(configFile, "")
}

// LoadForHostName will load the configuration from the named config file for specified host name,
// apply any overrides, and resolve relative directory locations.
func LoadForHostName(configFile, hostnameOverride string) (*Configuration, error) {
	f, err := DefaultFactory()
	if err != nil {
		return nil, errors.Trace(err)
	}
	config := new(Configuration)
	err = f.LoadConfigForHostName(configFile, hostnameOverride, config)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return config, nil
}
