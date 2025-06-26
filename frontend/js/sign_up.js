document.addEventListener('DOMContentLoaded', () => {
	const form = document.getElementById('sign_up_form');

	form.addEventListener('submit', async (e) => {
		e.preventDefault();

		const email = document.getElementById('email').value.trim();
		const password = document.getElementById('password').value;

		try {
			const response = await fetch('http://localhost:8080/api/users', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password })
			});

			if (response.ok) {
				const data = await response.json();

				localStorage.setItem('created_at', data.created_at);
				localStorage.setItem('updated_at', data.updated_at);
				localStorage.setItem('email', data.email);

				window.location.href = '/login.html';
			}
		} catch (err) {
			console.error('Account creation request failed:', err);
		}
	});
});
