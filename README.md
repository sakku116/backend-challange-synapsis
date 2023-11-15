# Backend Challenge Synapsis

## CLI Commands
- `seed-superuser`
    ```
    go run . seed-superuser
    ```
    or
    ```
    go run . seed-superuser superuser1  superuserpass
    ```
    - args (optional):
        args should be empty for default seed (superuser;superuser). Or args must be containing 2 strings for custom username, and passwords.


- `seed-data`
    used to seed all required data (minimum) for mandatory requirements.
    ```
    go run . seed-data
    ```

## API Documentation
see [http://localhost:8123/swagger/index.html](http://localhost:8123/swagger/index.html).

