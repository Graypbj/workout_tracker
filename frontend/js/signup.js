// frontend/js/signup.js
import { apiUrlBase } from './api.js'; // Use shared apiUrlBase

document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('signup-form');
    const errorMessageElement = document.getElementById('error-message');

    if (form) { // Ensure form exists
        form.addEventListener('submit', async (e) => {
            e.preventDefault();
            if(errorMessageElement) errorMessageElement.style.display = 'none';

            const email = document.getElementById('email').value.trim();
            const password = document.getElementById('password').value;

            if (!email || !password) {
                if(errorMessageElement){
                    errorMessageElement.textContent = 'Email and password are required.';
                    errorMessageElement.style.display = 'block';
                }
                return;
            }

            try {
                const response = await fetch(`${apiUrlBase}/users`, { // Use imported apiUrlBase
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ email, password })
                });

                if (response.status === 201) {
                    alert('Signup successful! Please login.');
                    window.location.href = 'index.html';
                } else {
                    const errorData = await response.json().catch(() => ({}));
                    let message = 'Signup failed. Please try again.';
                    if (errorData && errorData.error) {
                        message = errorData.error;
                    } else if (response.status === 400) {
                        message = 'Invalid data provided (e.g., email format, password too short).';
                    } else if (response.status === 409) {
                         message = 'This email is already registered.';
                    }
                    if(errorMessageElement){
                        errorMessageElement.textContent = message;
                        errorMessageElement.style.display = 'block';
                    }
                }
            } catch (err) {
                console.error('Signup request failed:', err);
                if(errorMessageElement){
                    errorMessageElement.textContent = 'An unexpected error occurred. Please try again.';
                    errorMessageElement.style.display = 'block';
                }
            }
        });
    }
});
