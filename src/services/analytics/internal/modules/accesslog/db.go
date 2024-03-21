package accesslog

import (
	"context"

	"github.com/pkg/errors"
	"github.com/varsotech/varsoapi/src/services/analytics/internal/ent"
)

/*
	Users can submit posts in either the main chart or the friends chart.
*/

func Create(ctx context.Context, uri, forwardedFor, ip, proto, host, port, server, id, userAgent string) error {
	err := ent.Database.AccessLog.Create().
		SetURI(uri).
		SetIP(ip).
		SetForwardedFor(forwardedFor).
		SetForwardedHost(host).
		SetForwardedPort(port).
		SetForwardedServer(server).
		SetForwardedProto(proto).
		SetRequestID(id).
		SetUserAgent(userAgent).
		Exec(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed creating access log in db")
	}

	return nil
}
