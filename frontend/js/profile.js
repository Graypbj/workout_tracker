// frontend/js/profile.js
import { apiUrlBase, getToken, handleUnauthorized, fetchWithAuth } from './api.js';

document.addEventListener('DOMContentLoaded', () => {
    // DOM Elements
    const updateProfileForm = document.getElementById('update-profile-form');
    const emailInput = document.getElementById('profile-email');
    const passwordInput = document.getElementById('profile-password');
    const messageElement = document.getElementById('profile-message'); // For success or error messages
    const logoutLink = document.getElementById('logout-link');

    // --- UI Helper ---
    const displayMessage = (message, isError = false) => {
        if (messageElement) {
            messageElement.textContent = message;
            messageElement.style.color = isError ? 'red' : 'green';
            messageElement.style.display = 'block';
        }
    };

    const clearMessage = () => {
        if (messageElement) {
            messageElement.textContent = '';
            messageElement.style.display = 'none';
        }
    };

    // --- Update Profile Functionality ---
    if (updateProfileForm) {
        updateProfileForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            clearMessage();

            const email = emailInput ? emailInput.value.trim() : '';
            const password = passwordInput ? passwordInput.value : '';

            if (!email || !password) {
                displayMessage('Both email and password are required to update your profile.', true);
                return;
            }

            if (password.length < 8) {
                displayMessage('Password must be at least 8 characters long.', true);
                return;
            }

            const profileData = { email, password };

            try {
                const response = await fetchWithAuth(`${apiUrlBase}/users`, {
                    method: 'PUT',
                    body: JSON.stringify(profileData)
                });

                // Check if response is OK (status 200-299)
                // For PUT, a 200 or 204 (if no content returned) can be success.
                // Let's assume 200 with updated user data, or 204 for success without data.
                if (response.ok) {
                    if (response.status === 204) { // No content returned, but successful
                        displayMessage('Profile updated successfully! Please note your new credentials.', false);
                    } else { // Assuming 200 OK with user data in response
                        const updatedUser = await response.json();
                        // Adjust according to actual backend response structure
                        const displayEmail = updatedUser?.User?.Email || updatedUser?.email || email;
                        displayMessage(`Profile updated successfully! Your email is now ${displayEmail}`, false);
                    }
                    if(updateProfileForm) updateProfileForm.reset();
                } else {
                    const errorData = await response.json().catch(() => ({ error: 'Failed to update profile due to a server error.' }));
                    displayMessage(errorData.error || 'Failed to update profile. Please try again.', true);
                }
            } catch (error) {
                console.error('Failed to update profile:', error);
                if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                     displayMessage('An error occurred while updating your profile. Please try again.', true);
                }
                // If it was 'Unauthorized' or 'No token found', fetchWithAuth already handled redirection.
            }
        });
    }

    // --- Logout ---
    if (logoutLink) {
        logoutLink.addEventListener('click', (e) => {
            e.preventDefault();
            handleUnauthorized();
        });
    }

    // --- Initial Load ---
    const init = () => {
        if (!getToken()) {
            handleUnauthorized();
            return;
        }
        if (emailInput) emailInput.focus();
    };

    if (window.location.pathname.endsWith('profile.html')) {
        init();
    }
});
