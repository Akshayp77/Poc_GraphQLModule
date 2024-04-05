package graph

import (
    "context"
    "fmt"
    model "GraphQLModule/graph/model"
)

// Resolver struct holds resolver functions
type Resolver struct{}

// Placeholder function to create a new organization
func createOrg(orgID string, orgName string, orgEmail string, orgAddress string, orgPhone string, orgType model.OrgType, createdAt string, updatedAt string) *model.Org {
    // Create a new organization with the provided data
    newOrg := &model.Org{
        OrgID:      orgID,
        OrgName:    orgName,
        OrgEmail:   orgEmail,
        OrgAddress: orgAddress,
        OrgPhone:   orgPhone,
        OrgType:    orgType,
        CreatedAt:  createdAt,
        UpdatedAt:  updatedAt,
    }
    return newOrg
}

// Resolver function for the createOrg mutation
func (r *Resolver) CreateOrg(ctx context.Context, args struct {
    OrgID      string
    OrgName    string
    OrgEmail   string
    OrgAddress string
    OrgPhone   string
    OrgType    model.OrgType
    CreatedAt  string
    UpdatedAt  string
}) (*model.Org, error) {
    // Your implementation logic goes here

    // Placeholder implementation
    newOrg := createOrg(args.OrgID, args.OrgName, args.OrgEmail, args.OrgAddress, args.OrgPhone, args.OrgType, args.CreatedAt, args.UpdatedAt)

    // Placeholder print statement
    fmt.Println("Organization Created:")
    fmt.Println("Organization ID:", newOrg.OrgID)
    fmt.Println("Organization Name:", newOrg.OrgName)
    fmt.Println("Organization Email:", newOrg.OrgEmail)
    fmt.Println("Organization Address:", newOrg.OrgAddress)
    fmt.Println("Organization Phone:", newOrg.OrgPhone)
    fmt.Println("Organization Type:", newOrg.OrgType)
    fmt.Println("Created At:", newOrg.CreatedAt)
    fmt.Println("Updated At:", newOrg.UpdatedAt)

    return newOrg, nil
}