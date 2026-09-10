package passkit

import (
	"testing"

	"github.com/PassKit/passkit-golang-grpc-sdk/io/members"
	"google.golang.org/protobuf/proto"
)

func TestMemberMetadataRemainsAProtobufMap(t *testing.T) {
	member := &members.Member{MetaData: map[string]string{"source": "quickstart"}}
	if _, err := proto.Marshal(member); err != nil {
		t.Fatalf("marshal member metadata: %v", err)
	}
}
