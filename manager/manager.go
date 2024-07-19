package manager

import (
	"broker_queues/common"
	"broker_queues/generated/message"

	"context"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

type Manager struct {
	rdb *redis.Client
	r   *mux.Router
	ctx context.Context
	ps  *redis.PubSub
}

func NewTaskHub() *Manager {
	th := &Manager{}

	th.rdb = redis.NewClient(&redis.Options{
		Addr: common.BrokerAddress,
	})
	th.ctx = context.Background()
	th.ps = th.rdb.Subscribe(th.ctx, common.BrokerMainChannel)

	go th.pubSubHandle()

	th.r = mux.NewRouter()

	th.r.HandleFunc("/", th.homeHandler)

	return th
}

func (th *Manager) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Test</h1>")
}

func (th *Manager) Run() error {
	http.Handle("/", th.r)
	err := http.ListenAndServe(":8080", th.r)
	return err
}

func (th *Manager) pubSubHandle() {
	defer th.ps.Close()

	for {
		msg, err := th.ps.ReceiveMessage(th.ctx)
		if err != nil {
			panic(err)
		}

		var message message.Message
		err = proto.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			continue
		}

		switch message.Data {
		case "exit":
			return
		}

		fmt.Println(&message)
	}
}
