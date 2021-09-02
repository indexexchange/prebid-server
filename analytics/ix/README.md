# IX Analytics Module

In order to use the ix analytics module, it needs to be configured by the host.

The pbs configuration file needs to be appened by the following block:

```yaml
analytics:
    ix:
      # Required properties
      enabled: true
      events: # Notification events that are enabled for logging
        win: true
        imp: false
      log_options: # Config expected by logger: github.com/chasex/glog
        file: "ix-prebid-server-analytics.log" # Required
        flag: 0 # Optional - Default: glog.LstdNull (0)
        level: 0 # Optional - Default: glog.Ldebug (0)
        mode: 3 # Optional - Default: glog.R_Day (3)
        maxsize: 0 # Optional - Default: 0

```