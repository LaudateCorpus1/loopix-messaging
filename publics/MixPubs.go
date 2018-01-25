/*
	Package publics implements struct for easy processing and storing of all public information
	of the network participants.
 */

package publics

import (
	"github.com/protobuf/proto"
)

func NewMixPubs(mixId, host, port string, pubKey []byte) MixPubs {
	mixPubs := MixPubs{Id: mixId, Host: host, Port: port, PubKey: pubKey}
	return mixPubs
}

func NewClientPubs(clientId, host, port string, pubKey []byte, providerInfo MixPubs) ClientPubs {
	client := ClientPubs{Id: clientId, Host: host, Port: port, PubKey: pubKey, Provider : &providerInfo}
	return client
}

func MixPubsToBytes(pubs MixPubs) ([]byte, error) {
	data, err := proto.Marshal(&pubs)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func MixPubsFromBytes(b []byte) (MixPubs, error) {
	var pubs MixPubs
	err := proto.Unmarshal(b, &pubs)
	if err != nil {
		return pubs, err
	}
	return pubs, nil
}

func ClientPubsToBytes(pubs ClientPubs) ([]byte, error) {
	data, err := proto.Marshal(&pubs)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func ClientPubsFromBytes(b []byte) (ClientPubs, error) {
	var pubs ClientPubs
	err := proto.Unmarshal(b, &pubs)
	if err != nil {
		return pubs, err
	}
	return pubs, nil
}


type E2EPath struct {
	IngressProvider MixPubs
	Mixes []MixPubs
	EgressProvider MixPubs
	Recipient ClientPubs
}

func (p *E2EPath) Len() int {
	return 3 + len(p.Mixes)
}