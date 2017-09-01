package util

import (
    "github.com/sirupsen/logrus"
    "database/sql"
)

var (
    // IsDebug False is release mode and true debug mode
    IsDebug = false

    // Log Logrus
    Log *logrus.Logger

    // Pgsql PostgreSQL connection
    Pgsql *sql.DB
)
