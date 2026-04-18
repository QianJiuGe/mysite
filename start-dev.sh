#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")" && pwd)"
BACK_DIR="$ROOT_DIR/backend"
FRONT_DIR="$ROOT_DIR/frontend"
RUNTIME_DIR="$ROOT_DIR/.runtime"

mkdir -p "$RUNTIME_DIR"

BACK_PID_FILE="$RUNTIME_DIR/backend.pid"
FRONT_PID_FILE="$RUNTIME_DIR/frontend.pid"
BACK_LOG="$RUNTIME_DIR/backend.log"
FRONT_LOG="$RUNTIME_DIR/frontend.log"

FRONT_BASE="http://localhost:5173"
BACK_BASE="http://localhost:8080"

GO_BIN_DEFAULT="$HOME/.local/go/bin/go"
if command -v go >/dev/null 2>&1; then
  GO_BIN="$(command -v go)"
elif [ -x "$GO_BIN_DEFAULT" ]; then
  GO_BIN="$GO_BIN_DEFAULT"
else
  echo "[ERROR] Go not found. Please install Go or place it at $GO_BIN_DEFAULT" >&2
  exit 1
fi

is_running() {
  local pid="$1"
  [ -n "$pid" ] && kill -0 "$pid" >/dev/null 2>&1
}

stop_if_running() {
  local name="$1"
  local pid_file="$2"

  if [ -f "$pid_file" ]; then
    local pid
    pid="$(cat "$pid_file" 2>/dev/null || true)"
    if is_running "$pid"; then
      echo "[INFO] Stopping stale $name process (pid=$pid)"
      kill "$pid" >/dev/null 2>&1 || true
      sleep 1
      if is_running "$pid"; then
        kill -9 "$pid" >/dev/null 2>&1 || true
      fi
    fi
    rm -f "$pid_file"
  fi
}

print_fail() {
  local step="$1"
  local msg="$2"
  echo "[ERROR] $step failed: $msg" >&2
  echo "[ERROR] Backend log: $BACK_LOG" >&2
  echo "[ERROR] Frontend log: $FRONT_LOG" >&2
  exit 1
}

echo "[1/4] Starting docker dependencies (MySQL + Redis)..."
if ! (cd "$BACK_DIR" && docker compose up -d >/dev/null); then
  print_fail "docker compose" "could not start MySQL/Redis"
fi

echo "[2/4] Starting backend..."
stop_if_running "backend" "$BACK_PID_FILE"
(
  cd "$BACK_DIR"
  nohup "$GO_BIN" run ./cmd/server -config ./configs/app.example.yaml >"$BACK_LOG" 2>&1 &
  echo $! >"$BACK_PID_FILE"
)

for _ in {1..20}; do
  if curl --noproxy '*' -fsS "$BACK_BASE/healthz" >/dev/null 2>&1; then
    break
  fi
  sleep 1
done

if ! curl --noproxy '*' -fsS "$BACK_BASE/healthz" >/dev/null 2>&1; then
  print_fail "backend" "health check did not pass on $BACK_BASE/healthz"
fi

echo "[3/4] Starting frontend..."
stop_if_running "frontend" "$FRONT_PID_FILE"
(
  cd "$FRONT_DIR"
  nohup npm run dev -- --host 0.0.0.0 --port 5173 >"$FRONT_LOG" 2>&1 &
  echo $! >"$FRONT_PID_FILE"
)

for _ in {1..20}; do
  if curl --noproxy '*' -fsSI "$FRONT_BASE" >/dev/null 2>&1; then
    break
  fi
  sleep 1
done

if ! curl --noproxy '*' -fsSI "$FRONT_BASE" >/dev/null 2>&1; then
  print_fail "frontend" "HTTP check failed on $FRONT_BASE"
fi

echo "[4/4] Done"

echo
echo "================ DEV SERVICES READY ================"
echo "Backend API:         $BACK_BASE"
echo "Backend health:      $BACK_BASE/healthz"
echo "Frontend root:       $FRONT_BASE"
echo "Login page:          $FRONT_BASE/login"
echo "Register page:       $FRONT_BASE/register"
echo "User home:           $FRONT_BASE/home"
echo "Admin console:       $FRONT_BASE/admin"
echo "Default admin:       admin / Admin@123456"
echo "----------------------------------------------------"
echo "Backend log:         $BACK_LOG"
echo "Frontend log:        $FRONT_LOG"
echo "Backend pid file:    $BACK_PID_FILE"
echo "Frontend pid file:   $FRONT_PID_FILE"
echo "===================================================="
