// frontend/js/login.js
import { apiUrlBase } from './api.js'; // Use shared apiUrlBase

document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('login-form');
    const errorMessageElement = document.getElementById('login-error-message');

    if (form) { // Ensure form exists before adding listener
        form.addEventListener('submit', async (e) => {
            e.preventDefault();
            if(errorMessageElement) errorMessageElement.style.display = 'none';

            const email = document.getElementById('email').value.trim();
            const password = document.getElementById('password').value;

            try {
                const response = await fetch(`${apiUrlBase}/login`, { // Use imported apiUrlBase
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ email, password })
                });

                if (response.ok) {
                    const data = await response.json();
                    localStorage.setItem('token', data.token);
                    localStorage.setItem('refreshToken', data.refresh_token);
                    window.location.href = '/workouts.html';
                } else {
                    const error = await response.json().catch(() => ({message: "Login failed"}));
                    let message = error.error || error.message || 'Login failed. Check credentials.'; // Adjusted to check error.error first
                    if(errorMessageElement) {
                        errorMessageElement.textContent = message;
                        errorMessageElement.style.display = 'block';
                    } else {
                        alert(message);
                    }
                }
            } catch (err) {
                console.error('Login request failed:', err);
                if(errorMessageElement) {
                    errorMessageElement.textContent = 'An error occurred. Please try again.';
                    errorMessageElement.style.display = 'block';
                } else {
                    alert('An error occurred. Please try again.');
                }
            }
        });
    }
});
