// internal/controller/controller_test.go
package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateVPNInstance(t *testing.T) {
	// Arrange
	controller := NewController()

	// Act
	vpn, err := controller.CreateVPNInstance("test-user")

	// Assert
	assert.NoError(t, err, "Expected no error when creating VPN instance")
	assert.NotNil(t, vpn, "Expected a non-nil VPN instance")
	assert.Equal(t, "test-user", vpn.UserID, "Expected UserID to match")
	assert.Equal(t, "active", vpn.Status, "Expected VPN status to be active")
	assert.NotEmpty(t, vpn.IP, "Expected VPN IP address to be non-empty")
}

func TestDeleteVPNInstance(t *testing.T) {
	controller := NewController()

	err := controller.DeleteVPNInstance("test-user")
	assert.NoError(t, err, "Expected no error when deleting VPN instance")
}