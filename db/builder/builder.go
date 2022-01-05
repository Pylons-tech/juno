package builder

import (
	"fmt"

	"github.com/pylons-tech/juno/db"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/pylons-tech/juno/config"
	"github.com/pylons-tech/juno/db/mongo"
	"github.com/pylons-tech/juno/db/postgresql"
)

// Builder represents a generic Builder implementation that build the proper database
// instance based on the configuration the user has specified
func Builder(cfg *config.Config, codec *codec.Codec) (db.Database, error) {
	switch cfg := cfg.DatabaseConfig.Config.(type) {
	case *config.MongoDBConfig:
		return mongo.Builder(cfg, codec)
	case *config.PostgreSQLConfig:
		return postgresql.Builder(cfg, codec)
	}

	return nil, fmt.Errorf("invalid config")
}
