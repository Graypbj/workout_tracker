// frontend/js/api.js
export const apiUrlBase = 'http://localhost:8080/api';

export const getToken = () => localStorage.getItem('token');

export const handleUnauthorized = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('refreshToken');
    window.location.href = '/index.html'; // Adjusted to root path for index.html
};

export const fetchWithAuth = async (url, options = {}) => {
    const token = getToken();
    // If no token, and it's not a public path like login/signup, handle unauthorized.
    // For simplicity, current public paths are handled by not calling fetchWithAuth.
    // This check is more for protected routes.
    if (!token) {
        handleUnauthorized();
        throw new Error('No token found, authorization required.');
    }

    const defaultHeaders = {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
    };

    options.headers = { ...defaultHeaders, ...options.headers };

    try {
        const response = await fetch(url, options);
        if (response.status === 401) {
            handleUnauthorized();
            throw new Error('Unauthorized');
        }
        // For DELETE requests, 204 No Content is a success response without a body.
        if (!response.ok && response.status !== 204) {
            const errorData = await response.json().catch(() => ({ error: 'Request failed with status: ' + response.status }));
            throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
        }
        return response;
    } catch (error) {
        console.error('API Fetch error:', error.message);
        throw error; // Re-throw to be caught by calling function
    }
};
