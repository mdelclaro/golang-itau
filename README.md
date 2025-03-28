# Golang Jokenpo (rock-paper-scissors)


## ⚡️ Running The Application

```bash
 make run
```

## 🧪 Test The Handler

```bash
 make test
```


## 📖 API Usage

### Endpoint: `/jokenpo/:option`

This endpoint allows you to play the Jokenpo (rock-paper-scissors) game against the computer.

- **Method**: `POST`
- **Path Parameter**: `:option` (integer)
  - `0` - Pedra (Rock)
  - `1` - Papel (Paper)
  - `2` - Tesoura (Scissors)

### Example Request

```bash
curl -X POST http://localhost:8080/jokenpo/0
```

### Example Response

```json
{
  "user": "pedra",
  "computer": "tesoura",
  "winner": "user"
}
```

- `user`: The option chosen by the user.
- `computer`: The option randomly chosen by the computer.
- `winner`: The winner of the game (`user`, `computer`, or `empate` for a tie).