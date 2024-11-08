package domain

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/iden3/go-iden3-core/v2/w3c"

	"github.com/polygonid/sh-id-platform/internal/common"
)

// DisplayMethodCoreDID is a type alias for w3c.DID
type DisplayMethodCoreDID w3c.DID

// DisplayMethod represents a display method
type DisplayMethod struct {
	ID        uuid.UUID
	Name      string
	URL       string
	IssuerDID DisplayMethodCoreDID
	IsDefault bool
}

// NewDisplayMethod creates a new display method
func NewDisplayMethod(id uuid.UUID, issuerDID w3c.DID, name, url string, isDefault bool) DisplayMethod {
	return DisplayMethod{
		ID:        id,
		IssuerDID: DisplayMethodCoreDID(issuerDID),
		Name:      name,
		URL:       url,
		IsDefault: isDefault,
	}
}

// IssuerCoreDID returns the issuer DID as a core DID
func (dm *DisplayMethod) IssuerCoreDID() *w3c.DID {
	return common.ToPointer(w3c.DID(dm.IssuerDID))
}

// Scan implements the sql.Scanner interface
func (dmdid *DisplayMethodCoreDID) Scan(value interface{}) error {
	didStr, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid value type, expected string")
	}
	did, err := w3c.ParseDID(didStr)
	if err != nil {
		return err
	}
	*dmdid = DisplayMethodCoreDID(*did)
	return nil
}
