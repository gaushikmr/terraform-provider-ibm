// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstanceReference p VM instance reference
// swagger:model PVMInstanceReference
type PVMInstanceReference struct {

	// (deprecated - replaced by networks) The list of addresses and their network information
	Addresses []*PVMInstanceNetwork `json:"addresses"`

	// Date/Time of PVM creation
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// fault
	Fault *PVMInstanceFault `json:"fault,omitempty"`

	// health
	Health *PVMInstanceHealth `json:"health,omitempty"`

	// Link to Cloud Instance resource
	// Required: true
	Href *string `json:"href"`

	// The ImageID used by the server
	// Required: true
	ImageID *string `json:"imageID"`

	// The list of addresses and their network information
	Networks []*PVMInstanceNetwork `json:"networks"`

	// The progress of an operation
	Progress float64 `json:"progress,omitempty"`

	// PCloud PVM Instance ID
	// Required: true
	PvmInstanceID *string `json:"pvmInstanceID"`

	// If this is an SAP pvm-instance the profile reference will link to the SAP profile
	SapProfile *SAPProfileReference `json:"sapProfile,omitempty"`

	// Name of the server
	// Required: true
	ServerName *string `json:"serverName"`

	// The pvm instance Software Licenses
	SoftwareLicenses *SoftwareLicenses `json:"softwareLicenses,omitempty"`

	// The status of the instance
	// Required: true
	Status *string `json:"status"`

	// System type used to host the instance
	SysType string `json:"sysType,omitempty"`

	// Date/Time of PVM last update
	// Format: date-time
	UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`
}

// Validate validates this p VM instance reference
func (m *PVMInstanceReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePvmInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSapProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceReference) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PVMInstanceReference) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceReference) validateFault(formats strfmt.Registry) error {

	if swag.IsZero(m.Fault) { // not required
		return nil
	}

	if m.Fault != nil {
		if err := m.Fault.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fault")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceReference) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceReference) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceReference) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageID", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceReference) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PVMInstanceReference) validatePvmInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("pvmInstanceID", "body", m.PvmInstanceID); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceReference) validateSapProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.SapProfile) { // not required
		return nil
	}

	if m.SapProfile != nil {
		if err := m.SapProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sapProfile")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceReference) validateServerName(formats strfmt.Registry) error {

	if err := validate.Required("serverName", "body", m.ServerName); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceReference) validateSoftwareLicenses(formats strfmt.Registry) error {

	if swag.IsZero(m.SoftwareLicenses) { // not required
		return nil
	}

	if m.SoftwareLicenses != nil {
		if err := m.SoftwareLicenses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceReference) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceReference) validateUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedDate", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceReference) UnmarshalBinary(b []byte) error {
	var res PVMInstanceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
