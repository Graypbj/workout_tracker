# API Documentation

## Authentication Endpoints

### 1. Login

*   **HTTP Method and Path:** `POST /api/login`
*   **Description:** Authenticates a user and returns a JWT access token and a refresh token.
*   **Request Body:**
    ```json
    {
        "email": "user@example.com",
        "password": "yourpassword"
    }
    ```
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "id": "uuid",
        "created_at": "timestamp",
        "updated_at": "timestamp",
        "email": "user@example.com",
        "token": "jwt_access_token",
        "refresh_token": "refresh_token_string"
    }
    ```
*   **Potential Error Responses:**
    *   `500 Internal Server Error`: "Couldn't decode parameters" - If the request body is malformed.
    *   `401 Unauthorized`: "Incorrect email or password" - If the email is not found or the password doesn't match.
    *   `500 Internal Server Error`: "Error generating token" - If JWT access token creation fails.
    *   `500 Internal Server Error`: "Error generating refresh token" - If refresh token creation fails.
    *   `500 Internal Server Error`: "Error generating refresh token db" - If storing the refresh token in the database fails.

### 2. Refresh Token

*   **HTTP Method and Path:** `POST /api/refresh`
*   **Description:** Takes a valid refresh token (sent in the Authorization header) and returns a new JWT access token.
*   **Request Headers:**
    *   `Authorization: Bearer <refresh_token>`
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "token": "new_jwt_access_token"
    }
    ```
*   **Potential Error Responses:**
    *   `400 Bad Request`: "Missing refresh token" - If the Authorization header is missing.
    *   `401 Unauthorized`: "Error finding refresh token in database" - If the provided refresh token is not found.
    *   `401 Unauthorized`: "Refresh token has been revoked" - If the refresh token has already been revoked.
    *   `401 Unauthorized`: "Refresh token has expired" - If the refresh token is past its expiry time.
    *   `500 Internal Server Error`: "Error retrieving user" - If the user associated with the token cannot be fetched.
    *   `500 Internal Server Error`: "Error creating access token" - If the new JWT access token creation fails.

### 3. Revoke Token

*   **HTTP Method and Path:** `POST /api/revoke`
*   **Description:** Revokes a refresh token (sent in the Authorization header).
*   **Request Headers:**
    *   `Authorization: Bearer <refresh_token>`
*   **Response Body (Success - 204 No Content):**
    *   Empty response.
*   **Potential Error Responses:**
    *   `400 Bad Request`: "No token sent" - If the Authorization header is missing.
    *   `500 Internal Server Error`: "Couldn't update token" - If an error occurs during the database update to revoke the token.

## User Management Endpoints

### 1. Create User

*   **HTTP Method and Path:** `POST /api/users`
*   **Description:** Creates a new user.
*   **Request Body:**
    ```json
    {
        "email": "newuser@example.com",
        "password": "newpassword"
    }
    ```
*   **Response Body (Success - 201 Created):**
    ```json
    {
        "id": "uuid",
        "created_at": "timestamp",
        "updated_at": "timestamp",
        "email": "newuser@example.com"
    }
    ```
*   **Potential Error Responses:**
    *   `500 Internal Server Error`: "Couldn't decode parameters" - If the request body is malformed.
    *   `500 Internal Server Error`: "Couldn't hash password" - If password hashing fails.
    *   `500 Internal Server Error`: "Couldn't create user" - If database user creation fails (e.g., email already exists).

### 2. Update User

*   **HTTP Method and Path:** `PUT /api/users`
*   **Description:** Updates an existing user's email or password. Requires authentication.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "email": "updateduser@example.com",
        "password": "updatedpassword"
    }
    ```
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "id": "uuid",
        "created_at": "timestamp",
        "updated_at": "timestamp",
        "email": "updateduser@example.com"
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT" - If the Authorization header is missing or malformed.
    *   `401 Unauthorized`: "Couldn't validate JWT" - If the JWT is invalid or expired.
    *   `500 Internal Server Error`: "Couldn't decode parameters" - If the request body is malformed.
    *   `500 Internal Server Error`: "Couldn't hash password" - If password hashing fails.
    *   `500 Internal Server Error`: "Couldn't update user" - If database user update fails.

### 3. Delete User

*   **HTTP Method and Path:** `DELETE /api/users`
*   **Description:** Deletes an existing user. Requires authentication.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Response Body (Success - 200 OK):**
    *   Returns an empty JSON object or null. *(The handler uses `var a any` which results in `null`)*
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT" - If the Authorization header is missing or malformed.
    *   `401 Unauthorized`: "Couldn't validate JWT" - If the JWT is invalid or expired.
    *   `500 Internal Server Error`: "Couldn't delete user" - If database user deletion fails.

## Exercise Management Endpoints

All exercise management endpoints require authentication.

### 1. Create Exercise

*   **HTTP Method and Path:** `POST /api/exercises`
*   **Description:** Creates a new exercise for the authenticated user.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "name": "Push Ups",
        "exercise_type": "Calisthenics"
    }
    ```
*   **Response Body (Success - 201 Created):**
    ```json
    {
        "id": "uuid",
        "user_id": "uuid",
        "name": "Push Ups",
        "exercise_type": "Calisthenics",
        "created_at": "timestamp",
        "updated_at": "timestamp"
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT" - If the Authorization header is missing or malformed.
    *   `401 Unauthorized`: "Couldn't validate JWT" - If the JWT is invalid or expired.
    *   `500 Internal Server Error`: "Couldn't decode parameters" - If the request body is malformed.
    *   `500 Internal Server Error`: "Couldn't create exercise" - If database exercise creation fails.

### 2. Retrieve Exercises

*   **HTTP Method and Path:** `GET /api/exercises`
*   **Description:** Retrieves all exercises for the authenticated user.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "exercises": [
            {
                "id": "uuid",
                "name": "Push Ups",
                "exercise_type": "Calisthenics",
                "created_at": "timestamp",
                "updated_at": "timestamp"
            },
            {
                "id": "uuid_2",
                "name": "Squats",
                "exercise_type": "Calisthenics",
                "created_at": "timestamp_2",
                "updated_at": "timestamp_2"
            }
        ]
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT" - If the Authorization header is missing or malformed.
    *   `401 Unauthorized`: "Couldn't validate JWT" - If the JWT is invalid or expired.
    *   `500 Internal Server Error`: "Couldn't retrieve exercises" - If there's an error fetching exercises from the database.

### 3. Update Exercise

*   **HTTP Method and Path:** `PUT /api/exercises`
*   **Description:** Updates an existing exercise for the authenticated user. The user must own the exercise.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "id": "exercise_uuid_to_update",
        "name": "Diamond Push Ups",
        "exercise_type": "Calisthenics"
    }
    ```
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "id": "exercise_uuid_to_update",
        "name": "Diamond Push Ups",
        "exercise_type": "Calisthenics",
        "created_at": "timestamp",
        "updated_at": "new_timestamp"
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT" - If the Authorization header is missing or malformed.
    *   `401 Unauthorized`: "Couldn't validate JWT" - If the JWT is invalid or expired.
    *   `500 Internal Server Error`: "Couldn't decode paramters" (Note: potential typo in server log, should be "parameters") - If the request body is malformed.
    *   `500 Internal Server Error`: "Couldn't update exercise" - If database exercise update fails (e.g., exercise not found, or user doesn't own it).

### 4. Delete Exercise

*   **HTTP Method and Path:** `DELETE /api/exercises`
*   **Description:** Deletes an existing exercise for the authenticated user. The user must own the exercise.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "id": "exercise_uuid_to_delete"
    }
    ```
*   **Response Body (Success - 200 OK):**
    *   Returns an empty JSON object or null. *(The handler uses `var a any` which results in `null`)*
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT" - If the Authorization header is missing or malformed.
    *   `401 Unauthorized`: "Couldn't validate JWT" - If the JWT is invalid or expired.
    *   `500 Internal Server Error`: "Couldn't decode parameters" - If the request body is malformed.
    *   `500 Internal Server Error`: "Couldn't delete exercise" - If database exercise deletion fails (e.g., exercise not found, or user doesn't own it).

## Workout Management Endpoints

All workout management endpoints require authentication.

### 1. Create Workout

*   **HTTP Method and Path:** `POST /api/workouts`
*   **Description:** Creates a new workout session for the authenticated user. `workout_date` is automatically set to the current time on creation.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "workout_type": "Full Body",
        "notes": "Felt good today, increased weights on bench."
    }
    ```
*   **Response Body (Success - 201 Created):**
    ```json
    {
        "id": "uuid",
        "user_id": "uuid", // Note: user_id is not actually returned by this handler as per the code
        "workout_date": "timestamp", // This is the creation time
        "workout_type": "Full Body",
        "notes": "Felt good today, increased weights on bench.",
        "created_at": "timestamp",
        "updated_at": "timestamp"
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't create workout"

### 2. Retrieve Workouts

*   **HTTP Method and Path:** `GET /api/workouts`
*   **Description:** Retrieves all workout sessions for the authenticated user.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "workouts": [
            {
                "id": "uuid1",
                "user_id": "uuid", // Note: user_id is not actually returned by this handler as per the code
                "workout_date": "timestamp1",
                "workout_type": "Full Body",
                "notes": "Session 1 notes",
                "created_at": "timestamp_c1",
                "updated_at": "timestamp_u1"
            },
            {
                "id": "uuid2",
                "user_id": "uuid", // Note: user_id is not actually returned by this handler as per the code
                "workout_date": "timestamp2",
                "workout_type": "Upper Body",
                "notes": "Session 2 notes",
                "created_at": "timestamp_c2",
                "updated_at": "timestamp_u2"
            }
        ]
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `500 Internal Server Error`: "Couldn't retrieve workouts"

### 3. Update Workout

*   **HTTP Method and Path:** `PUT /api/workouts`
*   **Description:** Updates an existing workout session for the authenticated user. The user must own the workout.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "id": "workout_uuid_to_update",
        "workout_date": "new_timestamp_for_workout", // Allows changing the recorded date of the workout
        "workout_type": "Lower Body",
        "notes": "Focused on squats and deadlifts."
    }
    ```
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "id": "workout_uuid_to_update",
        "user_id": "uuid", // Note: user_id is not actually returned by this handler as per the code
        "workout_date": "new_timestamp_for_workout",
        "workout_type": "Lower Body",
        "notes": "Focused on squats and deadlifts.",
        "created_at": "original_creation_timestamp",
        "updated_at": "new_update_timestamp"
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't update workout" (e.g., workout not found, or user doesn't own it)

### 4. Delete Workout

*   **HTTP Method and Path:** `DELETE /api/workouts`
*   **Description:** Deletes an existing workout session for the authenticated user. The user must own the workout.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "id": "workout_uuid_to_delete"
    }
    ```
*   **Response Body (Success - 200 OK):**
    *   Returns `null`.
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't delete workout" (e.g., workout not found, or user doesn't own it)

## Strength Training Session Endpoints

**Note on Authentication:** There are inconsistencies in authentication for these endpoints. Create and Update do not explicitly validate a JWT, which might be a security risk. Retrieve validates a JWT but doesn't confirm user ownership of the parent workout. Delete correctly validates the user against the session.

**Note on Update Logic:** The update operation (`PUT /api/strength_training_sessions`) appears to be missing the session ID in its parameters to `cfg.db.UpdateStrengthTrainingSessionByID`, which could lead to incorrect update behavior.

### 1. Create Strength Training Session

*   **HTTP Method and Path:** `POST /api/strength_training_sessions`
*   **Description:** Creates a new strength training session, typically associated with a workout and an exercise.
*   **Authentication:** *Not explicitly performed in the handler. This might be a security concern.*
*   **Request Body:**
    ```json
    {
        "workout_id": "workout_uuid",
        "exercise_id": "exercise_uuid",
        "notes": "Performed 3 sets of 10 reps."
    }
    ```
*   **Response Body (Success - 201 Created):**
    ```json
    {
        "strength_training_session": {
            "id": "new_session_uuid",
            "workout_id": "workout_uuid",
            "exercise_id": "exercise_uuid",
            "notes": "Performed 3 sets of 10 reps.",
            "created_at": "timestamp",
            "updated_at": "timestamp"
        }
    }
    ```
*   **Potential Error Responses:**
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't create strength training session"

### 2. Retrieve Strength Training Sessions by Workout

*   **HTTP Method and Path:** `GET /api/strength_training_sessions?workout_id=<workout_uuid>`
*   **Description:** Retrieves all strength training sessions associated with a specific `workout_id`.
*   **Authentication:** Requires JWT. *However, it does not explicitly check if the authenticated user owns the specified `workout_id`.*
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Query Parameters:**
    *   `workout_id` (required): The UUID of the workout.
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "strength_training_sessions": [
            {
                "id": "session_uuid_1",
                "workout_id": "workout_uuid",
                "exercise_id": "exercise_uuid_1",
                "notes": "Notes for session 1",
                "created_at": "timestamp",
                "updated_at": "timestamp"
            },
            {
                "id": "session_uuid_2",
                "workout_id": "workout_uuid",
                "exercise_id": "exercise_uuid_2",
                "notes": "Notes for session 2",
                "created_at": "timestamp",
                "updated_at": "timestamp"
            }
        ]
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `400 Bad Request`: "Missing workout_id query parameter"
    *   `400 Bad Request`: "Invalid workout_id format"
    *   `500 Internal Server Error`: "Couldn't retrieve strength training sessions"

### 3. Update Strength Training Session

*   **HTTP Method and Path:** `PUT /api/strength_training_sessions`
*   **Description:** Updates an existing strength training session.
*   **Authentication:** *Not explicitly performed in the handler. This might be a security concern.*
*   **Note:** The current implementation seems to lack a specific session ID for the update, potentially leading to incorrect behavior. It uses `WorkoutID` and `ExerciseID` which might not be unique identifiers for a session.
*   **Request Body:**
    ```json
    {
        // "id": "session_uuid_to_update", // This ID seems to be missing from the db params
        "workout_id": "workout_uuid", // Used in db params
        "exercise_id": "new_exercise_uuid", // Used in db params
        "notes": "Updated notes for the session."
    }
    ```
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "strength_training_session": {
            "id": "session_uuid_that_was_updated", // This ID comes from the DB response
            "workout_id": "workout_uuid",
            "exercise_id": "new_exercise_uuid",
            "notes": "Updated notes for the session.",
            "created_at": "timestamp",
            "updated_at": "new_timestamp"
        }
    }
    ```
*   **Potential Error Responses:**
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't update strength training session"

### 4. Delete Strength Training Session

*   **HTTP Method and Path:** `DELETE /api/strength_training_sessions`
*   **Description:** Deletes a specific strength training session. Requires authentication, and the user must be associated with the session for deletion.
*   **Authentication:** Requires JWT. The `userID` from the token is used in the delete database query.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "id": "session_uuid_to_delete"
    }
    ```
*   **Response Body (Success - 200 OK):**
    *   Returns `null`.
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't delete strength training session" (e.g., session not found, or user mismatch)

## Strength Training Set Endpoints

All strength training set endpoints require authentication and operate within the context of the authenticated user.

### 1. Create Strength Training Set

*   **HTTP Method and Path:** `POST /api/strength_training_sets`
*   **Description:** Creates a new set for a strength training session.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "session_id": "strength_session_uuid",
        "set_number": 1,
        "reps": 10,
        "weight": "100kg"
    }
    ```
*   **Response Body (Success - 201 Created):**
    ```json
    {
        "strength_training_set": {
            "id": "new_set_uuid",
            // "user_id": "user_uuid", // Not included in response
            "session_id": "strength_session_uuid",
            "set_number": 1,
            "reps": 10,
            "weight": "100kg",
            "created_at": "timestamp",
            "updated_at": "timestamp"
        }
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't create strength training set" (e.g., session_id does not exist or does not belong to user)

### 2. Retrieve Strength Training Sets by Session

*   **HTTP Method and Path:** `GET /api/strength_training_sets?session_id=<session_uuid>`
*   **Description:** Retrieves all sets for a specific strength training session belonging to the authenticated user.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Query Parameters:**
    *   `session_id` (required): The UUID of the strength training session.
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "strength_training_sets": [
            {
                "id": "set_uuid_1",
                // "user_id": "user_uuid", // Not included in response
                "session_id": "session_uuid",
                "set_number": 1,
                "reps": 10,
                "weight": "100kg",
                "created_at": "timestamp",
                "updated_at": "timestamp"
            },
            {
                "id": "set_uuid_2",
                // "user_id": "user_uuid", // Not included in response
                "session_id": "session_uuid",
                "set_number": 2,
                "reps": 8,
                "weight": "105kg",
                "created_at": "timestamp",
                "updated_at": "timestamp"
            }
        ]
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `400 Bad Request`: "Missing session_id query parameter"
    *   `400 Bad Request`: "Invalid workout_id format" (Typo in error message, should be "Invalid session_id format")
    *   `500 Internal Server Error`: "Couldn't retrieve sets"

### 3. Update Strength Training Set

*   **HTTP Method and Path:** `PUT /api/strength_training_sets`
*   **Description:** Updates an existing strength training set. The set must belong to the authenticated user.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "id": "set_uuid_to_update",
        "set_number": 1,
        "reps": 12,
        "weight": "95kg"
    }
    ```
*   **Response Body (Success - 200 OK):**
    ```json
    {
        "strength_training_set": {
            "id": "set_uuid_to_update",
            // "user_id": "user_uuid", // Not included in response
            "session_id": "original_session_uuid", // session_id is part of the response struct but not updatable here
            "set_number": 1,
            "reps": 12,
            "weight": "95kg",
            "created_at": "original_timestamp",
            "updated_at": "new_timestamp"
        }
    }
    ```
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't update strength training set" (e.g., set not found or user mismatch)

### 4. Delete Strength Training Set

*   **HTTP Method and Path:** `DELETE /api/strength_training_sets`
*   **Description:** Deletes a specific strength training set. The set must belong to the authenticated user.
*   **Request Headers:**
    *   `Authorization: Bearer <jwt_access_token>`
*   **Request Body:**
    ```json
    {
        "id": "set_uuid_to_delete"
    }
    ```
*   **Response Body (Success - 200 OK):**
    *   Returns `null`.
*   **Potential Error Responses:**
    *   `401 Unauthorized`: "Couldn't find JWT"
    *   `401 Unauthorized`: "Couldn't validate JWT"
    *   `500 Internal Server Error`: "Couldn't decode parameters"
    *   `500 Internal Server Error`: "Couldn't delete strength training set" (e.g., set not found or user mismatch)
