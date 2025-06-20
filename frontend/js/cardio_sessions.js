// frontend/js/cardio_sessions.js
import { apiUrlBase, getToken, handleUnauthorized, fetchWithAuth } from './api.js';

document.addEventListener('DOMContentLoaded', () => {
    // DOM Elements
    const manageSessionForm = document.getElementById('manage-cardio-session-form');
    const sessionIdInput = document.getElementById('cardio-session-id');
    const exerciseNameInput = document.getElementById('cardio-exercise-name');
    const durationInput = document.getElementById('cardio-duration');
    const distanceInput = document.getElementById('cardio-distance');
    const sessionDateInput = document.getElementById('cardio-session-date');
    const notesInput = document.getElementById('cardio-notes');
    const sessionsListElement = document.getElementById('cardio-sessions-list');
    const cancelEditBtn = document.getElementById('cancel-edit-btn');
    const errorMessageElement = document.getElementById('cardio-session-error-message');
    const logoutLink = document.getElementById('logout-link');

    // --- UI Helper ---
    const displayErrorMessage = (message) => {
        if (errorMessageElement) {
            errorMessageElement.textContent = message;
            errorMessageElement.style.display = 'block';
        }
    };
    const clearErrorMessage = () => {
        if (errorMessageElement) {
            errorMessageElement.textContent = '';
            errorMessageElement.style.display = 'none';
        }
    };
    const resetForm = () => {
        if (manageSessionForm) manageSessionForm.reset();
        if (sessionIdInput) sessionIdInput.value = '';
        if (cancelEditBtn) cancelEditBtn.style.display = 'none';
        if (exerciseNameInput) exerciseNameInput.focus();
        if (manageSessionForm) manageSessionForm.querySelector('button[type="submit"]').textContent = 'Save Session';
        clearErrorMessage();
    };

    // --- CRUD Functions ---
    const displayCardioSessions = async () => {
        clearErrorMessage();
        if (!getToken()) { handleUnauthorized(); return; }
        if (!sessionsListElement) return;

        try {
            const response = await fetchWithAuth(`${apiUrlBase}/cardio_training_sessions`);
            let sessions = [];
            if (response.status !== 204) { // Check for content
                const data = await response.json();
                sessions = data.cardio_training_sessions || [];
            }


            sessionsListElement.innerHTML = '';

            if (sessions.length === 0) {
                sessionsListElement.innerHTML = '<p>No cardio sessions found. Add some!</p>';
                return;
            }

            sessions.forEach(session => {
                const sessionElement = document.createElement('div');
                sessionElement.classList.add('session-item');
                sessionElement.innerHTML = `
                    <h4>${session.exercise_name}</h4>
                    <p><strong>Date:</strong> ${new Date(session.session_date).toLocaleDateString()}</p>
                    <p><strong>Duration:</strong> ${session.duration_minutes} minutes</p>
                    ${session.distance_km ? `<p><strong>Distance:</strong> ${session.distance_km} km</p>` : ''}
                    ${session.notes ? `<p><strong>Notes:</strong> ${session.notes}</p>` : ''}
                    <div>
                        <button class="edit-btn" data-id="${session.id}">Edit</button>
                        <button class="delete-btn" data-id="${session.id}">Delete</button>
                    </div>
                `;
                sessionsListElement.appendChild(sessionElement);

                const editBtn = sessionElement.querySelector('.edit-btn');
                if (editBtn) editBtn.addEventListener('click', () => populateFormForEdit(session));

                const deleteBtn = sessionElement.querySelector('.delete-btn');
                if (deleteBtn) deleteBtn.addEventListener('click', () => deleteCardioSession(session.id));
            });
        } catch (error) {
            console.error('Failed to display cardio sessions:', error);
            if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                displayErrorMessage(`Failed to load cardio sessions: ${error.message}`);
            }
        }
    };

    const populateFormForEdit = (session) => {
        clearErrorMessage();
        if(sessionIdInput) sessionIdInput.value = session.id;
        if(exerciseNameInput) exerciseNameInput.value = session.exercise_name;
        if(durationInput) durationInput.value = session.duration_minutes;
        if(distanceInput) distanceInput.value = session.distance_km || '';
        if(sessionDateInput) sessionDateInput.value = session.session_date.split('T')[0];
        if(notesInput) notesInput.value = session.notes || '';

        if(manageSessionForm) manageSessionForm.querySelector('button[type="submit"]').textContent = 'Update Session';
        if(cancelEditBtn) cancelEditBtn.style.display = 'inline-block';
        if(exerciseNameInput) exerciseNameInput.focus();
    };

    if(manageSessionForm) {
        manageSessionForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            clearErrorMessage();

            const id = sessionIdInput ? sessionIdInput.value : null;
            const exercise_name = exerciseNameInput ? exerciseNameInput.value.trim() : '';
            const duration_minutes = durationInput ? parseInt(durationInput.value) : NaN;
            const distance_km = distanceInput && distanceInput.value ? parseFloat(distanceInput.value) : null;
            const session_date = sessionDateInput ? sessionDateInput.value : null;
            const notes = notesInput ? notesInput.value.trim() : '';

            if (!exercise_name || isNaN(duration_minutes) || !session_date) {
                displayErrorMessage('Exercise name, duration, and date are required.');
                return;
            }
            if (duration_minutes <=0) {
                displayErrorMessage('Duration must be greater than 0.');
                return;
            }
            if (distance_km !== null && distance_km < 0) {
                displayErrorMessage('Distance cannot be negative.');
                return;
            }

            const sessionData = { exercise_name, duration_minutes, distance_km, session_date, notes };
            const method = id ? 'PUT' : 'POST';
            if (id) {
                sessionData.id = id;
            }

            try {
                await fetchWithAuth(`${apiUrlBase}/cardio_training_sessions`, {
                    method: method,
                    body: JSON.stringify(sessionData)
                });
                resetForm();
                // manageSessionForm.querySelector('button[type="submit"]').textContent = 'Save Session'; // Handled by resetForm
                await displayCardioSessions();
            } catch (error) {
                console.error('Failed to save cardio session:', error);
                if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                    displayErrorMessage(`Failed to save session: ${error.message}`);
                }
            }
        });
    }

    const deleteCardioSession = async (id) => {
        if (!confirm('Are you sure you want to delete this cardio session?')) {
            return;
        }
        clearErrorMessage();
        try {
            await fetchWithAuth(`${apiUrlBase}/cardio_training_sessions`, {
                method: 'DELETE',
                body: JSON.stringify({ id: id })
            });
            await displayCardioSessions();
        } catch (error) {
            console.error('Failed to delete cardio session:', error);
            if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                displayErrorMessage(`Failed to delete session: ${error.message}`);
            }
        }
    };

    if(cancelEditBtn) {
        cancelEditBtn.addEventListener('click', () => {
            resetForm();
            // manageSessionForm.querySelector('button[type="submit"]').textContent = 'Save Session'; // Handled by resetForm
        });
    }

    if (logoutLink) {
        logoutLink.addEventListener('click', (e) => {
            e.preventDefault();
            handleUnauthorized();
        });
    }

    // Initial Load
    const init = async () => {
        if (!getToken()) {
            handleUnauthorized();
            return;
        }
        await displayCardioSessions();
    };

    if (window.location.pathname.endsWith('cardio_sessions.html')) {
        init();
    }
});
