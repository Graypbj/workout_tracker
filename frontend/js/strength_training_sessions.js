// frontend/js/strength_training_sessions.js
import { apiUrlBase, getToken, handleUnauthorized, fetchWithAuth } from './api.js';

document.addEventListener('DOMContentLoaded', () => {
    // DOM Elements
    const createSessionForm = document.getElementById('create-session-form');
    const sessionDateInput = document.getElementById('session-date');
    const sessionNotesInput = document.getElementById('session-notes');
    const sessionsListElement = document.getElementById('sessions-list');
    const sessionErrorMessageElement = document.getElementById('session-error-message');
    const logoutLink = document.getElementById('logout-link');

    // Set Modal Elements
    const setModal = document.getElementById('manage-set-modal');
    const setModalTitle = document.getElementById('set-modal-title');
    const manageSetForm = document.getElementById('manage-set-form');
    const setIdInput = document.getElementById('set-id');
    const setSessionIdInput = document.getElementById('set-session-id');
    const setExerciseSelect = document.getElementById('set-exercise');
    const setRepsInput = document.getElementById('set-reps');
    const setWeightInput = document.getElementById('set-weight');
    const setDurationInput = document.getElementById('set-duration');
    const setDistanceInput = document.getElementById('set-distance');
    const setNotesFormInput = document.getElementById('set-notes-form'); // Corrected ID
    const cancelSetModalBtn = document.getElementById('cancel-set-modal');
    const setErrorMessageElement = document.getElementById('set-error-message');

    let exercisesCache = []; // To store fetched exercises

    // --- Utility Functions ---
    const displaySessionError = (message) => {
        if (sessionErrorMessageElement) {
            sessionErrorMessageElement.textContent = message;
            sessionErrorMessageElement.style.display = 'block';
        }
    };
    const clearSessionError = () => {
        if (sessionErrorMessageElement) {
            sessionErrorMessageElement.textContent = '';
            sessionErrorMessageElement.style.display = 'none';
        }
    };
    const displaySetError = (message) => {
        if (setErrorMessageElement) {
            setErrorMessageElement.textContent = message;
            setErrorMessageElement.style.display = 'block';
        }
    };
    const clearSetError = () => {
        if (setErrorMessageElement) {
            setErrorMessageElement.textContent = '';
            setErrorMessageElement.style.display = 'none';
        }
    };

    // --- Load Exercises for Select Dropdown ---
    const loadExercises = async () => {
        if (!getToken()) { handleUnauthorized(); return; }
        try {
            const response = await fetchWithAuth(`${apiUrlBase}/exercises`);
            if (response.status === 204) { // No content
                exercisesCache = [];
            } else {
                const data = await response.json();
                exercisesCache = data.exercises || [];
            }
            populateExerciseSelect(exercisesCache);
        } catch (error) {
            console.error('Failed to load exercises:', error);
            displaySetError('Failed to load exercises for set management.');
        }
    };

    const populateExerciseSelect = (exercises) => {
        if (!setExerciseSelect) return;
        setExerciseSelect.innerHTML = '<option value="">Select Exercise</option>';
        exercises.forEach(ex => {
            const option = document.createElement('option');
            option.value = ex.id;
            option.textContent = ex.name;
            setExerciseSelect.appendChild(option);
        });
    };

    // --- Session CRUD ---
    const displaySessions = async () => {
        clearSessionError();
        if (!getToken()) { handleUnauthorized(); return; }
        if (!sessionsListElement) return;

        try {
            const response = await fetchWithAuth(`${apiUrlBase}/strength_training_sessions`);
            let sessions = [];
            if (response.status !== 204) { // Check if there is content
                const data = await response.json();
                sessions = data.strength_training_sessions || [];
            }


            sessionsListElement.innerHTML = '';
            if (sessions.length === 0) {
                sessionsListElement.innerHTML = '<p>No strength training sessions found. Add one!</p>';
                return;
            }

            for (const session of sessions) {
                const sessionElement = document.createElement('div');
                sessionElement.classList.add('session');
                sessionElement.innerHTML = `
                    <div class="session-header">
                        <div>
                            <h4>Session on: ${new Date(session.session_date).toLocaleDateString()}</h4>
                            <p>Notes: ${session.notes || 'N/A'}</p>
                        </div>
                        <div class="session-actions">
                            <button class="edit-session-btn" data-id="${session.id}">Edit Session</button>
                            <button class="delete-session-btn" data-id="${session.id}">Delete Session</button>
                            <button class="add-set-btn" data-session-id="${session.id}">Add Set</button>
                        </div>
                    </div>
                    <div class="sets-list" id="sets-for-session-${session.id}">Loading sets...</div>
                `;
                sessionsListElement.appendChild(sessionElement);

                sessionElement.querySelector('.delete-session-btn').addEventListener('click', () => deleteSession(session.id));
                sessionElement.querySelector('.add-set-btn').addEventListener('click', () => openSetModal(session.id));
                sessionElement.querySelector('.edit-session-btn').addEventListener('click', () => populateSessionFormForEdit(session));

                await displaySetsForSession(session.id);
            }
        } catch (error) {
            console.error('Failed to display sessions:', error);
            if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                displaySessionError(`Error loading sessions: ${error.message}`);
            }
        }
    };

    const populateSessionFormForEdit = (session) => {
        console.log("Edit session clicked:", session);
        if (sessionDateInput) sessionDateInput.value = session.session_date.split('T')[0];
        if (sessionNotesInput) sessionNotesInput.value = session.notes;

        // This is a simplified approach. A real edit would likely involve:
        // 1. A hidden input field in the form for the session ID.
        //    e.g., let idField = document.getElementById('session-editing-id'); if (!idField) { idField = document.createElement('input'); idField.type = 'hidden'; idField.id = 'session-editing-id'; createSessionForm.appendChild(idField); } idField.value = session.id;
        // 2. Changing the submit button text to "Update Session".
        //    e.g., if (createSessionForm) createSessionForm.querySelector('button[type="submit"]').textContent = "Update Session";
        // 3. The form submission handler would check for this ID and use PUT method.
        alert("Session editing UI not fully implemented. Details logged to console. You can adapt the 'Add New Session' form for editing by populating fields and changing the submit action to a PUT request.");
    };

    if (createSessionForm) {
        createSessionForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            clearSessionError();
            const date = sessionDateInput ? sessionDateInput.value : null;
            const notes = sessionNotesInput ? sessionNotesInput.value.trim() : '';

            if (!date) {
                displaySessionError('Session date is required.');
                return;
            }

            // Example for handling edit (see populateSessionFormForEdit for more context)
            // const editingId = document.getElementById('session-editing-id')?.value;
            // const method = editingId ? 'PUT' : 'POST';
            // const requestBody = { session_date: date, notes };
            // if (editingId) requestBody.id = editingId;

            try {
                await fetchWithAuth(`${apiUrlBase}/strength_training_sessions`, {
                    method: 'POST', // Simplified: only POST for now
                    body: JSON.stringify({ session_date: date, notes })
                });
                createSessionForm.reset();
                // if (document.getElementById('session-editing-id')) document.getElementById('session-editing-id').value = ''; // Clear editing ID
                // createSessionForm.querySelector('button[type="submit"]').textContent = "Add Session";
                await displaySessions();
            } catch (error) {
                console.error('Failed to save session:', error);
                if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                     displaySessionError(`Error saving session: ${error.message}`);
                }
            }
        });
    }

    const deleteSession = async (sessionId) => {
        if (!confirm('Are you sure you want to delete this entire session and all its sets?')) return;
        clearSessionError();
        try {
            await fetchWithAuth(`${apiUrlBase}/strength_training_sessions`, {
                method: 'DELETE',
                body: JSON.stringify({ id: sessionId })
            });
            await displaySessions();
        } catch (error) {
            console.error('Failed to delete session:', error);
             if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                displaySessionError(`Error deleting session: ${error.message}`);
            }
        }
    };

    // --- Set CRUD ---
    const displaySetsForSession = async (sessionId) => {
        const setsListContainer = document.getElementById(`sets-for-session-${sessionId}`);
        if (!setsListContainer) return;

        try {
            const response = await fetchWithAuth(`${apiUrlBase}/strength_training_sets?session_id=${sessionId}`);
            let sets = [];
            if(response.status !== 204) { // Check for content
                const data = await response.json();
                sets = data.strength_training_sets || [];
            }

            setsListContainer.innerHTML = '';
            if (sets.length === 0) {
                setsListContainer.innerHTML = '<p>No sets recorded for this session yet.</p>';
                return;
            }

            sets.forEach(set => {
                const exerciseName = exercisesCache.find(ex => ex.id === set.exercise_id)?.name || 'Unknown Exercise';
                const setElement = document.createElement('div');
                setElement.classList.add('set-item');
                setElement.innerHTML = `
                    <p><strong>${exerciseName}</strong></p>
                    <p>Reps: ${set.reps}, Weight: ${set.weight} kg</p>
                    ${set.duration_seconds ? `<p>Duration: ${set.duration_seconds}s</p>` : ''}
                    ${set.distance_meters ? `<p>Distance: ${set.distance_meters}m</p>` : ''}
                    ${set.notes ? `<p>Notes: ${set.notes}</p>` : ''}
                    <div class="set-actions">
                        <button class="edit-set-btn" data-set-id="${set.id}" data-session-id="${sessionId}">Edit Set</button>
                        <button class="delete-set-btn" data-set-id="${set.id}" data-session-id="${sessionId}">Delete Set</button>
                    </div>
                `;
                setsListContainer.appendChild(setElement);
                setElement.querySelector('.edit-set-btn').addEventListener('click', () => openSetModal(sessionId, set));
                setElement.querySelector('.delete-set-btn').addEventListener('click', () => deleteSet(set.id, sessionId));
            });
        } catch (error) {
            console.error(`Failed to display sets for session ${sessionId}:`, error);
            if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                 setsListContainer.innerHTML = '<p class="error-text">Could not load sets.</p>';
            }
        }
    };

    const openSetModal = (sessionId, setToEdit = null) => {
        clearSetError();
        if(manageSetForm) manageSetForm.reset();
        if(setSessionIdInput) setSessionIdInput.value = sessionId;

        if (setToEdit) {
            if(setModalTitle) setModalTitle.textContent = 'Edit Set';
            if(setIdInput) setIdInput.value = setToEdit.id;
            if(setExerciseSelect) setExerciseSelect.value = setToEdit.exercise_id;
            if(setRepsInput) setRepsInput.value = setToEdit.reps;
            if(setWeightInput) setWeightInput.value = setToEdit.weight;
            if(setDurationInput) setDurationInput.value = setToEdit.duration_seconds || '';
            if(setDistanceInput) setDistanceInput.value = setToEdit.distance_meters || '';
            if(setNotesFormInput) setNotesFormInput.value = setToEdit.notes || '';
        } else {
            if(setModalTitle) setModalTitle.textContent = 'Add Set';
            if(setIdInput) setIdInput.value = '';
        }
        if(setModal) setModal.style.display = 'flex';
    };

    const closeSetModal = () => {
        if(setModal) setModal.style.display = 'none';
        clearSetError();
    };

    if(cancelSetModalBtn) cancelSetModalBtn.addEventListener('click', closeSetModal);
    if(setModal) {
        setModal.addEventListener('click', (e) => {
            if (e.target === setModal) closeSetModal();
        });
    }

    if(manageSetForm) {
        manageSetForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            clearSetError();

            const setId = setIdInput ? setIdInput.value : null;
            const sessionId = setSessionIdInput ? setSessionIdInput.value : null;
            const exerciseId = setExerciseSelect ? setExerciseSelect.value : null;
            const reps = setRepsInput ? parseInt(setRepsInput.value) : NaN;
            const weight = setWeightInput ? parseFloat(setWeightInput.value) : NaN;
            const duration = setDurationInput && setDurationInput.value ? parseInt(setDurationInput.value) : null;
            const distance = setDistanceInput && setDistanceInput.value ? parseInt(setDistanceInput.value) : null;
            const notes = setNotesFormInput ? setNotesFormInput.value.trim() : '';

            if (!sessionId || !exerciseId || isNaN(reps) || isNaN(weight)) {
                displaySetError('Session context, exercise, reps, and weight are required.');
                return;
            }

            const setData = {
                session_id: sessionId,
                exercise_id: exerciseId,
                reps: reps,
                weight: weight,
                notes: notes
            };
            if(duration !== null && !isNaN(duration)) setData.duration_seconds = duration;
            if(distance !== null && !isNaN(distance)) setData.distance_meters = distance;

            const isEditing = !!setId;
            const method = isEditing ? 'PUT' : 'POST';
            if (isEditing) {
                setData.id = setId;
            }

            try {
                await fetchWithAuth(`${apiUrlBase}/strength_training_sets`, {
                    method: method,
                    body: JSON.stringify(setData)
                });
                closeSetModal();
                await displaySetsForSession(sessionId);
            } catch (error) {
                console.error('Failed to save set:', error);
                if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                    displaySetError(`Error saving set: ${error.message}`);
                }
            }
        });
    }

    const deleteSet = async (setId, sessionId) => {
        if (!confirm('Are you sure you want to delete this set?')) return;
        clearSetError();
        try {
            await fetchWithAuth(`${apiUrlBase}/strength_training_sets`, {
                method: 'DELETE',
                body: JSON.stringify({ id: setId })
            });
            await displaySetsForSession(sessionId);
        } catch (error) {
            console.error('Failed to delete set:', error);
            if (error.message !== 'Unauthorized' && error.message !== 'No token found, authorization required.') {
                alert(`Failed to delete set: ${error.message}`);
            }
        }
    };

    // --- Logout ---
    if (logoutLink) {
        logoutLink.addEventListener('click', (e) => {
            e.preventDefault();
            handleUnauthorized();
        });
    }

    // --- Initial Load ---
    const init = async () => {
        if (!getToken()) {
            handleUnauthorized();
            return;
        }
        await loadExercises();
        await displaySessions();
    };

    if (window.location.pathname.endsWith('strength_training_sessions.html')) {
        init();
    }
});
