<script lang="ts">
  import { onMount } from 'svelte';
  import type { Player } from '$lib/types/player';
  import PlayersList from './_components/PlayersList.svelte';
  import Leaderboard from './_components/Leaderboard.svelte';
  import WinnerAnnouncement from './_components/WinnerAnnouncement.svelte';
  import Instructions from './_components/Instructions.svelte';

  let players = $state<Player[]>([]);
  let newPlayerName = $state('');
  let gameStarted = $state(false);
  let timeLeft = $state(5);
  let gameEnded = $state(false);
  let currentPlayerIndex = $state(0);
  let playerStarted = $state(false);
  let inCooling = $state(false);
  let coolingTime = $state(3);
  let timer: number | undefined;

  onMount(() => {
    // Load players from localStorage
    const savedPlayers = localStorage.getItem('fingerTapPlayers');
    if (savedPlayers) {
      try {
        players = JSON.parse(savedPlayers);
      } catch (e) {
        console.error('Failed to load saved players:', e);
      }
    }
  });

  function addPlayer() {
    if (newPlayerName.trim() && players.length < 10) {
      players = [...players, { name: newPlayerName.trim(), taps: 0 }];
      newPlayerName = '';
      savePlayersToStorage();
    }
  }

  function removePlayer(index: number) {
    players = players.filter((_, i) => i !== index);
    savePlayersToStorage();
  }

  function clearAllPlayers() {
    if (confirm('Are you sure you want to clear all players?')) {
      players = [];
      localStorage.removeItem('fingerTapPlayers');
    }
  }

  function savePlayersToStorage() {
    localStorage.setItem('fingerTapPlayers', JSON.stringify(players));
  }

  function handleTap(index: number) {
    if (!gameStarted || gameEnded || index !== currentPlayerIndex || inCooling)
      return;

    // Start timer on first tap
    if (!playerStarted) {
      playerStarted = true;
      startTimer();
    }

    // Update taps and trigger reactivity
    players = players.map((p, i) =>
      i === index ? { ...p, taps: p.taps + 1 } : p,
    );
  }

  function startGame() {
    if (players.length < 2) return;

    gameStarted = true;
    gameEnded = false;
    timeLeft = 5;
    currentPlayerIndex = 0;
    playerStarted = false;
    inCooling = false;

    // Reset all taps
    players = players.map((p) => ({ ...p, taps: 0 }));
  }

  function startTimer() {
    // Clear any existing timer
    if (timer) clearInterval(timer);

    // Start countdown
    timer = window.setInterval(() => {
      timeLeft--;
      if (timeLeft <= 0) {
        // Clear timer and start cooling period
        clearInterval(timer);
        startCoolingPeriod();
      }
    }, 1000);
  }

  function startCoolingPeriod() {
    inCooling = true;
    coolingTime = 3;
    playerStarted = false;

    timer = window.setInterval(() => {
      coolingTime--;
      if (coolingTime <= 0) {
        clearInterval(timer);
        inCooling = false;

        // Move to next player or end game
        if (currentPlayerIndex < players.length - 1) {
          currentPlayerIndex++;
          timeLeft = 5;
        } else {
          endGame();
        }
      }
    }, 1000);
  }

  function endGame() {
    gameStarted = false;
    gameEnded = true;
    if (timer) {
      clearInterval(timer);
    }
  }

  function resetGame() {
    gameEnded = false;
    players = players.map((p) => ({ ...p, taps: 0 }));
  }

  function handleKeyPress(e: KeyboardEvent) {
    if (e.key === 'Enter') {
      addPlayer();
    }
  }

  $effect(() => {
    return () => {
      if (timer) {
        clearInterval(timer);
      }
    };
  });

  // Get winner(s)
  const maxTaps = $derived.by(() => {
    if (players.length === 0) return 0;
    return Math.max(...players.map((p) => p.taps), 0);
  });

  const winners = $derived.by(() => {
    if (!gameEnded || players.length === 0) return [];
    const max = Math.max(...players.map((p) => p.taps), 0);
    return players.filter((p) => p.taps === max && max > 0);
  });

  const sortedPlayers = $derived.by(() => {
    return [...players].sort((a, b) => b.taps - a.taps);
  });
</script>

<svelte:head>
  <title>Finger Tap Challenge | Maejic Code</title>
</svelte:head>

<div
  class="min-h-screen bg-linear-to-br from-orange-500 via-red-500 to-pink-600 py-12 px-4 sm:px-6 lg:px-8 pt-30"
>
  <div class="max-w-6xl mx-auto">
    <!-- Header -->
    <div class="text-center mb-8">
      <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold text-white mb-4">
        👆 Finger Tap Challenge
      </h1>
      <p class="text-lg md:text-xl text-white/90">
        Each player gets 5 seconds to tap!
      </p>
    </div>

    {#if !gameStarted && !gameEnded}
      <!-- Setup Phase -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
        <!-- Add Players Section -->
        <div class="bg-white/10 backdrop-blur-lg rounded-2xl p-6">
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-2xl font-bold text-white">
              Players ({players.length})
            </h2>
            {#if players.length > 0}
              <button
                onclick={clearAllPlayers}
                class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-lg font-semibold text-sm transition-colors"
              >
                Clear All
              </button>
            {/if}
          </div>

          <!-- Add player input -->
          <div class="mb-6">
            <div class="flex gap-2">
              <input
                type="text"
                bind:value={newPlayerName}
                onkeypress={handleKeyPress}
                placeholder="Enter player name..."
                maxlength="20"
                class="flex-1 px-4 py-3 rounded-lg bg-white/20 text-white placeholder-white/60 border-2 border-white/30 focus:border-yellow-400 focus:outline-none"
              />
              <button
                onclick={addPlayer}
                disabled={!newPlayerName.trim() || players.length >= 10}
                class="bg-green-500 hover:bg-green-600 text-white px-6 py-3 rounded-lg font-semibold disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                Add
              </button>
            </div>
            <p class="text-white/70 text-sm mt-2">Maximum 10 players</p>
          </div>

          <!-- Players list -->
          <PlayersList {players} {removePlayer} />
        </div>

        <!-- Instructions -->
        <Instructions />
      </div>

      <!-- Start Button -->
      <div class="text-center px-4 py-6">
        <button
          onclick={startGame}
          disabled={players.length < 2}
          class="bg-linear-to-r from-yellow-400 to-orange-500 hover:from-yellow-500 hover:to-orange-600 text-white text-xl font-bold px-12 py-6 rounded-2xl shadow-2xl disabled:opacity-50 disabled:cursor-not-allowed transition-all transform hover:scale-105 disabled:hover:scale-100 cursor-pointer"
        >
          {#if players.length < 2}
            Add at least 2 players to start
          {:else}
            🎮 Start Game
          {/if}
        </button>
      </div>

      <!-- Back to Games (Setup Phase) -->
      <div class="text-center mt-6">
        <a
          href="/games"
          class="inline-flex items-center gap-2 bg-white/10 hover:bg-white/20 backdrop-blur-lg text-white px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105"
        >
          <span>←</span>
          <span>Back to Games</span>
        </a>
      </div>
    {/if}

    {#if gameStarted && !gameEnded}
      <!-- Game Phase -->
      <div class="space-y-6">
        <!-- Timer -->
        <div class="text-center">
          <div
            class="inline-flex flex-col items-center bg-white/20 backdrop-blur-lg rounded-3xl px-6 py-8 shadow-2xl"
          >
            {#if inCooling}
              <div class="text-white/80 text-xl mb-2">Get Ready!</div>
              <div class="text-white/60 text-lg mb-4">
                {#if currentPlayerIndex < players.length - 1}
                  Next: {players[currentPlayerIndex + 1]?.name}
                {:else}
                  Final Results Coming...
                {/if}
              </div>
              <div class="text-8xl font-bold text-white animate-pulse">
                {coolingTime}
              </div>
            {:else}
              <div class="text-white/80 text-lg mb-1">
                {players[currentPlayerIndex]?.name}'s Turn
              </div>
              {#if playerStarted}
                <div class="text-white/80 text-xl mb-2">Time Remaining</div>
                <div class="text-8xl font-bold text-white animate-pulse">
                  {timeLeft}
                </div>
              {:else}
                <div class="text-white/80 text-xl mb-2">Ready?</div>
                <div
                  class="text-6xl font-bold text-yellow-400 animate-bounce py-4"
                >
                  Tap to Start!
                </div>
              {/if}
            {/if}
          </div>
        </div>

        <!-- Tap Buttons -->
        {#if !inCooling}
          <div class="flex justify-center mt-8">
            {#each players as player, i}
              {#if i === currentPlayerIndex}
                <button
                  onclick={() => handleTap(i)}
                  class="bg-white/10 backdrop-blur-lg hover:bg-white/20 rounded-3xl p-8 md:p-12 transition-all transform active:scale-95 hover:scale-105 shadow-2xl ring-4 ring-yellow-400 animate-pulse w-full max-w-md"
                >
                  <div class="text-center">
                    <div class="text-6xl md:text-8xl mb-6">👆</div>
                    <div class="text-white font-bold text-3xl md:text-4xl mb-4">
                      {player.name}
                    </div>
                    <div
                      class="text-8xl md:text-9xl font-bold text-yellow-400 mb-2"
                    >
                      {player.taps}
                    </div>
                    <div
                      class="text-white/80 text-xl md:text-2xl font-bold mb-4"
                    >
                      {#if playerStarted}
                        TAP NOW!
                      {:else}
                        TAP TO START!
                      {/if}
                    </div>
                    <div class="text-white/60 text-sm">
                      Player {i + 1} of {players.length}
                    </div>
                  </div>
                </button>
              {/if}
            {/each}
          </div>
        {/if}
      </div>
    {/if}

    {#if gameEnded}
      <!-- Results Phase -->
      <div class="space-y-6">
        <!-- Winner Announcement -->
        <WinnerAnnouncement {winners} {maxTaps} />

        <!-- Leaderboard -->
        <Leaderboard {sortedPlayers} />

        <!-- Action Buttons -->
        <div class="flex gap-4 justify-center flex-wrap mb-6">
          <button
            onclick={startGame}
            class="bg-linear-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 text-white text-lg font-bold px-8 py-4 rounded-xl shadow-xl transition-all transform hover:scale-105"
          >
            🔄 Play Again
          </button>
          <button
            onclick={resetGame}
            class="bg-white/20 hover:bg-white/30 backdrop-blur-lg text-white text-lg font-bold px-8 py-4 rounded-xl transition-all transform hover:scale-105"
          >
            📝 Change Players
          </button>
        </div>

        <!-- Back to Games -->
        <div class="text-center">
          <a
            href="/games"
            class="inline-flex items-center gap-2 bg-white/10 hover:bg-white/20 backdrop-blur-lg text-white px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105"
          >
            <span>←</span>
            <span>Back to Games</span>
          </a>
        </div>
      </div>
    {/if}
  </div>
</div>

<style>
  @keyframes scaleIn {
    from {
      transform: scale(0.5);
      opacity: 0;
    }
    to {
      transform: scale(1);
      opacity: 1;
    }
  }
</style>
