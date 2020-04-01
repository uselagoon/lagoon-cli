package schema

import (
	"encoding/json"
	"fmt"

	"github.com/amazeeio/lagoon-cli/pkg/api"
	"github.com/google/uuid"
)

// AddGroupInput is based on the input to addGroup.
type AddGroupInput struct {
	Name        string      `json:"name"`
	ParentGroup *GroupInput `json:"parentGroup,omitempty"`
}

// Group provides for unmarshalling the groups contained with a Project.
type Group struct {
	AddGroupInput
	ID      *uuid.UUID `json:"id,omitempty"`
	Members []struct {
		User User          `json:"user"`
		Role api.GroupRole `json:"role"`
	} `json:"members,omitempty"`
}

// GroupConfig embeds AddGroupInput as well as a list of members.
type GroupConfig struct {
	AddGroupInput
	Users []UserRoleConfig `json:"users,omitempty"`
}

// GroupInput is based on the Lagoon API type.
type GroupInput ProjectInput

// UserGroupRoleInput is based on the Lagoon API type.
type UserGroupRoleInput struct {
	UserEmail string        `json:"userEmail"`
	GroupName string        `json:"groupName"`
	GroupRole api.GroupRole `json:"groupRole"`
}

// UserRoleConfig stores a user/role config within a group.
type UserRoleConfig struct {
	Email string        `json:"email"`
	Role  api.GroupRole `json:"role"`
}

// AddBillingGroupInput is based on the input to addBillingGroup.
type AddBillingGroupInput struct {
	Name            string   `json:"name"`
	Currency        Currency `json:"currency"`
	BillingSoftware string   `json:"billingSoftware,omitempty"`
}

// BillingGroup provides for unmarshalling the groups contained with a Project.
type BillingGroup struct {
	AddBillingGroupInput
	ID *uuid.UUID `json:"id,omitempty"`
}

// Groups represents possible Lagoon group types.
// These are unmarshalled from a projectByName query response.
type Groups struct {
	Groups        []Group
	BillingGroups []BillingGroup
}

// UnmarshalJSON unmashals a quoted json string to the Notification values.
func (g *Groups) UnmarshalJSON(b []byte) error {
	var gArrayRaw []map[string]json.RawMessage
	if err := json.Unmarshal(b, &gArrayRaw); err != nil {
		return err
	}
	possibleKeys := []string{"__typename", "name", "currency", "billingSoftware"}
	var value string
	for _, groupMapRaw := range gArrayRaw {
		if len(groupMapRaw) == 0 {
			// Unsupported group type returns an empty map... even when the unknown
			// type it represents is not requested! (╯°□°）╯︵ ┻━┻
			// This happens when the lagoon API being targeted is actually a higher
			// version than configured.
			continue
		}
		gMap := map[string]string{}
		for _, k := range possibleKeys {
			rawMsg, ok := groupMapRaw[k]
			if !ok {
				continue // key missing
			}
			if err := json.Unmarshal(rawMsg, &value); err != nil {
				return err
			}
			gMap[k] = value
		}

		switch gMap["__typename"] {
		case "Group":
			group := Group{
				AddGroupInput: AddGroupInput{
					Name: gMap["name"],
				},
			}
			err := json.Unmarshal(groupMapRaw["members"], &group.Members)
			if err != nil {
				return err
			}
			g.Groups = append(g.Groups, group)
		case "BillingGroup":
			g.BillingGroups = append(g.BillingGroups,
				BillingGroup{
					AddBillingGroupInput: AddBillingGroupInput{
						Name:            gMap["name"],
						Currency:        Currency(gMap["currency"]),
						BillingSoftware: gMap["billingSoftware"],
					},
				})
		case "":
			return fmt.Errorf(`missing key "__typename" in group response`)
		default:
			return fmt.Errorf("unknown group type: %s", gMap["__typename"])
		}
	}
	return nil
}
