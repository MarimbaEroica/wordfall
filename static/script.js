document.addEventListener('DOMContentLoaded', () => {
    const socket = new WebSocket(`ws://${window.location.host}/ws`);

    let board = [];
    let selectedTiles = [];
    let score = 0;
    let timeLeft = 180;

    // WebSocket event handlers
    socket.onopen = function(event) {
        console.log('Connected to server');
    };

    socket.onmessage = function(event) {
        const msg = JSON.parse(event.data);
        switch (msg.Type) {
            case 'boardUpdate':
                const payload = msg.Payload;
                board = payload.Board;
                score = payload.Score;
                timeLeft = payload.TimeLeft;
                updateScore();
                updateTimer();
                renderBoard();
                break;
            case 'invalidWord':
                alert('Invalid word or path!');
                clearSelection();
                break;
            // Handle other message types as needed
        }
    };

    function updateScore() {
        document.getElementById('score').textContent = score;
    }

    function updateTimer() {
        document.getElementById('timer').textContent = timeLeft;
        if (timeLeft <= 0) {
            alert('Time is up! Your score is ' + score);
        }
    }

    // Render the game board
    function renderBoard() {
        const gameBoard = document.getElementById('game-board');
        gameBoard.innerHTML = '';
        for (let col = 0; col < board.length; col++) {
            const columnDiv = document.createElement('div');
            columnDiv.classList.add('column');
            const column = board[col];
            for (let row = 0; row < column.length; row++) {
                const tileDiv = document.createElement('div');
                tileDiv.classList.add('tile');
                const tileLetter = column[row];
                tileDiv.textContent = tileLetter || '';
                tileDiv.dataset.col = col;
                tileDiv.dataset.row = row;
                if (row >= 5) {
                    tileDiv.classList.add('hidden-tile');
                } else {
                    tileDiv.addEventListener('click', selectTile);
                }
                tileDiv.style.transform = 'translateY(-100%)';
                columnDiv.appendChild(tileDiv);
                // Animate the tile falling
                setTimeout(() => {
                    tileDiv.style.transition = 'transform 0.5s';
                    tileDiv.style.transform = 'translateY(0)';
                }, 50 * row); // Delay each tile slightly for a cascading effect
            }
            gameBoard.appendChild(columnDiv);
        }
    }

    // Handle tile selection
    function selectTile(e) {
        const tileDiv = e.target;
        tileDiv.classList.toggle('selected');
        const col = parseInt(tileDiv.dataset.col);
        const row = parseInt(tileDiv.dataset.row);
        const tileIndex = selectedTiles.findIndex(tile => tile.col === col && tile.row === row);
        if (tileIndex > -1) {
            selectedTiles.splice(tileIndex, 1);
        } else {
            selectedTiles.push({ col, row });
        }
    }

    // Submit word
    function submitWord() {
        const message = {
            Type: 'wordSubmission',
            Payload: {
                SelectedTiles: selectedTiles
            }
        };
        socket.send(JSON.stringify(message));
        clearSelection();
    }

    // Remove tiles manually
    function removeTilesManually() {
        const message = {
            Type: 'manualRemoval',
            Payload: {
                SelectedTiles: selectedTiles
            }
        };
        socket.send(JSON.stringify(message));
        clearSelection();
    }

    function clearSelection() {
        selectedTiles = [];
        const selectedDivs = document.querySelectorAll('.selected');
        selectedDivs.forEach(div => div.classList.remove('selected'));
    }

    // Event listeners
    document.getElementById('submit-word').addEventListener('click', submitWord);
    document.getElementById('remove-tiles').addEventListener('click', removeTilesManually);
});
