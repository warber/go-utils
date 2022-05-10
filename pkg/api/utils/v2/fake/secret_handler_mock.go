// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package utils_mock

import (
	"context"
	"github.com/keptn/go-utils/pkg/api/models"
	"github.com/keptn/go-utils/pkg/api/utils/v2"
	"sync"
)

// SecretsInterfaceMock is a mock implementation of v2.SecretsInterface.
//
// 	func TestSomethingThatUsesSecretsInterface(t *testing.T) {
//
// 		// make and configure a mocked v2.SecretsInterface
// 		mockedSecretsInterface := &SecretsInterfaceMock{
// 			CreateSecretFunc: func(ctx context.Context, secret models.Secret, opts v2.SecretsCreateSecretOptions) error {
// 				panic("mock out the CreateSecret method")
// 			},
// 			DeleteSecretFunc: func(ctx context.Context, secretName string, secretScope string, opts v2.SecretsDeleteSecretOptions) error {
// 				panic("mock out the DeleteSecret method")
// 			},
// 			GetSecretsFunc: func(ctx context.Context, opts v2.SecretsGetSecretsOptions) (*models.GetSecretsResponse, error) {
// 				panic("mock out the GetSecrets method")
// 			},
// 			UpdateSecretFunc: func(ctx context.Context, secret models.Secret, opts v2.SecretsUpdateSecretOptions) error {
// 				panic("mock out the UpdateSecret method")
// 			},
// 		}
//
// 		// use mockedSecretsInterface in code that requires v2.SecretsInterface
// 		// and then make assertions.
//
// 	}
type SecretsInterfaceMock struct {
	// CreateSecretFunc mocks the CreateSecret method.
	CreateSecretFunc func(ctx context.Context, secret models.Secret, opts v2.SecretsCreateSecretOptions) error

	// DeleteSecretFunc mocks the DeleteSecret method.
	DeleteSecretFunc func(ctx context.Context, secretName string, secretScope string, opts v2.SecretsDeleteSecretOptions) error

	// GetSecretsFunc mocks the GetSecrets method.
	GetSecretsFunc func(ctx context.Context, opts v2.SecretsGetSecretsOptions) (*models.GetSecretsResponse, error)

	// UpdateSecretFunc mocks the UpdateSecret method.
	UpdateSecretFunc func(ctx context.Context, secret models.Secret, opts v2.SecretsUpdateSecretOptions) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateSecret holds details about calls to the CreateSecret method.
		CreateSecret []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Secret is the secret argument value.
			Secret models.Secret
			// Opts is the opts argument value.
			Opts v2.SecretsCreateSecretOptions
		}
		// DeleteSecret holds details about calls to the DeleteSecret method.
		DeleteSecret []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SecretName is the secretName argument value.
			SecretName string
			// SecretScope is the secretScope argument value.
			SecretScope string
			// Opts is the opts argument value.
			Opts v2.SecretsDeleteSecretOptions
		}
		// GetSecrets holds details about calls to the GetSecrets method.
		GetSecrets []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts v2.SecretsGetSecretsOptions
		}
		// UpdateSecret holds details about calls to the UpdateSecret method.
		UpdateSecret []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Secret is the secret argument value.
			Secret models.Secret
			// Opts is the opts argument value.
			Opts v2.SecretsUpdateSecretOptions
		}
	}
	lockCreateSecret sync.RWMutex
	lockDeleteSecret sync.RWMutex
	lockGetSecrets   sync.RWMutex
	lockUpdateSecret sync.RWMutex
}

// CreateSecret calls CreateSecretFunc.
func (mock *SecretsInterfaceMock) CreateSecret(ctx context.Context, secret models.Secret, opts v2.SecretsCreateSecretOptions) error {
	if mock.CreateSecretFunc == nil {
		panic("SecretsInterfaceMock.CreateSecretFunc: method is nil but SecretsInterface.CreateSecret was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Secret models.Secret
		Opts   v2.SecretsCreateSecretOptions
	}{
		Ctx:    ctx,
		Secret: secret,
		Opts:   opts,
	}
	mock.lockCreateSecret.Lock()
	mock.calls.CreateSecret = append(mock.calls.CreateSecret, callInfo)
	mock.lockCreateSecret.Unlock()
	return mock.CreateSecretFunc(ctx, secret, opts)
}

// CreateSecretCalls gets all the calls that were made to CreateSecret.
// Check the length with:
//     len(mockedSecretsInterface.CreateSecretCalls())
func (mock *SecretsInterfaceMock) CreateSecretCalls() []struct {
	Ctx    context.Context
	Secret models.Secret
	Opts   v2.SecretsCreateSecretOptions
} {
	var calls []struct {
		Ctx    context.Context
		Secret models.Secret
		Opts   v2.SecretsCreateSecretOptions
	}
	mock.lockCreateSecret.RLock()
	calls = mock.calls.CreateSecret
	mock.lockCreateSecret.RUnlock()
	return calls
}

// DeleteSecret calls DeleteSecretFunc.
func (mock *SecretsInterfaceMock) DeleteSecret(ctx context.Context, secretName string, secretScope string, opts v2.SecretsDeleteSecretOptions) error {
	if mock.DeleteSecretFunc == nil {
		panic("SecretsInterfaceMock.DeleteSecretFunc: method is nil but SecretsInterface.DeleteSecret was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		SecretName  string
		SecretScope string
		Opts        v2.SecretsDeleteSecretOptions
	}{
		Ctx:         ctx,
		SecretName:  secretName,
		SecretScope: secretScope,
		Opts:        opts,
	}
	mock.lockDeleteSecret.Lock()
	mock.calls.DeleteSecret = append(mock.calls.DeleteSecret, callInfo)
	mock.lockDeleteSecret.Unlock()
	return mock.DeleteSecretFunc(ctx, secretName, secretScope, opts)
}

// DeleteSecretCalls gets all the calls that were made to DeleteSecret.
// Check the length with:
//     len(mockedSecretsInterface.DeleteSecretCalls())
func (mock *SecretsInterfaceMock) DeleteSecretCalls() []struct {
	Ctx         context.Context
	SecretName  string
	SecretScope string
	Opts        v2.SecretsDeleteSecretOptions
} {
	var calls []struct {
		Ctx         context.Context
		SecretName  string
		SecretScope string
		Opts        v2.SecretsDeleteSecretOptions
	}
	mock.lockDeleteSecret.RLock()
	calls = mock.calls.DeleteSecret
	mock.lockDeleteSecret.RUnlock()
	return calls
}

// GetSecrets calls GetSecretsFunc.
func (mock *SecretsInterfaceMock) GetSecrets(ctx context.Context, opts v2.SecretsGetSecretsOptions) (*models.GetSecretsResponse, error) {
	if mock.GetSecretsFunc == nil {
		panic("SecretsInterfaceMock.GetSecretsFunc: method is nil but SecretsInterface.GetSecrets was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts v2.SecretsGetSecretsOptions
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockGetSecrets.Lock()
	mock.calls.GetSecrets = append(mock.calls.GetSecrets, callInfo)
	mock.lockGetSecrets.Unlock()
	return mock.GetSecretsFunc(ctx, opts)
}

// GetSecretsCalls gets all the calls that were made to GetSecrets.
// Check the length with:
//     len(mockedSecretsInterface.GetSecretsCalls())
func (mock *SecretsInterfaceMock) GetSecretsCalls() []struct {
	Ctx  context.Context
	Opts v2.SecretsGetSecretsOptions
} {
	var calls []struct {
		Ctx  context.Context
		Opts v2.SecretsGetSecretsOptions
	}
	mock.lockGetSecrets.RLock()
	calls = mock.calls.GetSecrets
	mock.lockGetSecrets.RUnlock()
	return calls
}

// UpdateSecret calls UpdateSecretFunc.
func (mock *SecretsInterfaceMock) UpdateSecret(ctx context.Context, secret models.Secret, opts v2.SecretsUpdateSecretOptions) error {
	if mock.UpdateSecretFunc == nil {
		panic("SecretsInterfaceMock.UpdateSecretFunc: method is nil but SecretsInterface.UpdateSecret was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Secret models.Secret
		Opts   v2.SecretsUpdateSecretOptions
	}{
		Ctx:    ctx,
		Secret: secret,
		Opts:   opts,
	}
	mock.lockUpdateSecret.Lock()
	mock.calls.UpdateSecret = append(mock.calls.UpdateSecret, callInfo)
	mock.lockUpdateSecret.Unlock()
	return mock.UpdateSecretFunc(ctx, secret, opts)
}

// UpdateSecretCalls gets all the calls that were made to UpdateSecret.
// Check the length with:
//     len(mockedSecretsInterface.UpdateSecretCalls())
func (mock *SecretsInterfaceMock) UpdateSecretCalls() []struct {
	Ctx    context.Context
	Secret models.Secret
	Opts   v2.SecretsUpdateSecretOptions
} {
	var calls []struct {
		Ctx    context.Context
		Secret models.Secret
		Opts   v2.SecretsUpdateSecretOptions
	}
	mock.lockUpdateSecret.RLock()
	calls = mock.calls.UpdateSecret
	mock.lockUpdateSecret.RUnlock()
	return calls
}
