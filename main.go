package main

import (
    "github.com/sirupsen/logrus"
    "github.com/heroku/rollrus"
    "gitlab.com/avegao/iot-thermostat/util"
    "github.com/evalphobia/logrus_sentry"
)

func initLogger() {
    util.IsDebug  = util.GetBoolParameter("DEBUG", util.IsDebug )

    logLevel := logrus.InfoLevel

    if util.IsDebug  {
        logLevel = logrus.DebugLevel
    }

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetLevel(logLevel)

    log := logrus.New()
    log.SetLevel(logLevel)

    initSentry(log)
    initRollbar(log)

    util.Log = log
}

func initSentry(log *logrus.Logger) {
    dsn := util.GetStringParameter("sentry_dsn", "")

    tags := make(map[string]string)
    tags["environment"] = "develop"

    if !util.IsDebug {
        tags["environment"] = "release"
    }

    hook, err := logrus_sentry.NewWithTagsSentryHook(dsn, tags, []logrus.Level{
        logrus.PanicLevel,
        logrus.FatalLevel,
        logrus.ErrorLevel,
        logrus.WarnLevel,
    })

    if nil == err {
        log.Hooks.Add(hook)
    }
}

func initRollbar(log *logrus.Logger) {
    token := util.GetStringParameter("rollbar_token", "")
    environment := "develop"

    if !util.IsDebug {
        environment = "release"
    }

    hook := rollrus.NewHookForLevels(token, environment, []logrus.Level{
        logrus.PanicLevel,
        logrus.FatalLevel,
        logrus.ErrorLevel,
        logrus.WarnLevel,
    })

    log.Hooks.Add(hook)
}

func main() {
    initLogger()

    util.Log.Infoln("Hello World!")
}
