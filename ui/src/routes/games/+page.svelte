<script lang="ts">
  import { goto } from '$app/navigation';

  interface Game {
    title: string;
    description: string;
    icon: string;
    path: string;
    color: string;
  }

  const games: Game[] = [
    {
      title: 'Roulette Wheel',
      description: 'Spin the wheel to randomly select from your list of names',
      icon: '🎡',
      path: '/games/roulette',
      color: 'from-pink-500 to-purple-600',
    },
    {
      title: 'Finger Tap Challenge',
      description:
        'Fast-paced tapping competition - see who can tap the fastest in 5 seconds!',
      icon: '👆',
      path: '/games/finger-tap',
      color: 'from-orange-500 to-red-600',
    },
    {
      title: 'Coming Soon',
      description: 'More exciting games are on the way!',
      icon: '🎮',
      path: '#',
      color: 'from-gray-400 to-gray-600',
    },
  ];

  function navigateToGame(path: string) {
    if (path !== '#') {
      goto(path);
    }
  }
</script>

<div
  class="min-h-screen bg-linear-to-br from-purple-900 via-blue-900 to-indigo-900 pb-12 pt-30 px-4 sm:px-6 lg:px-8"
>
  <div class="max-w-7xl mx-auto">
    <!-- Header -->
    <div class="text-center mb-12">
      <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold text-white mb-4">
        🎮 Game Collection
      </h1>
      <p class="text-lg md:text-xl text-gray-300">
        Choose your favorite game and have fun!
      </p>
    </div>

    <!-- Games Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 md:gap-8">
      {#each games as game}
        <button
          onclick={() => navigateToGame(game.path)}
          disabled={game.path === '#'}
          class="group relative bg-white/10 backdrop-blur-lg rounded-2xl p-6 md:p-8 hover:bg-white/20 transition-all duration-300 hover:scale-105 hover:shadow-2xl disabled:hover:scale-100 disabled:cursor-not-allowed disabled:opacity-60"
        >
          <!-- Gradient overlay -->
          <div
            class="absolute inset-0 bg-linear-to-br {game.color} opacity-20 rounded-2xl group-hover:opacity-30 transition-opacity"
          ></div>

          <!-- Content -->
          <div class="relative z-10">
            <!-- Icon -->
            <div
              class="text-6xl md:text-7xl mb-4 transform group-hover:scale-110 transition-transform"
            >
              {game.icon}
            </div>

            <!-- Title -->
            <h2 class="text-2xl md:text-3xl font-bold text-white mb-3">
              {game.title}
            </h2>

            <!-- Description -->
            <p class="text-gray-300 text-sm md:text-base mb-4">
              {game.description}
            </p>

            <!-- Play button -->
            {#if game.path !== '#'}
              <div
                class="inline-flex items-center gap-2 bg-white/20 px-4 py-2 rounded-lg font-semibold text-white group-hover:bg-white/30 transition-colors"
              >
                <span>Play Now</span>
                <span
                  class="transform group-hover:translate-x-1 transition-transform"
                  >→</span
                >
              </div>
            {:else}
              <div
                class="inline-flex items-center gap-2 bg-white/10 px-4 py-2 rounded-lg font-semibold text-gray-400"
              >
                <span>Coming Soon</span>
              </div>
            {/if}
          </div>
        </button>
      {/each}
    </div>

    <!-- Back to Home -->
    <div class="text-center mt-12">
      <a
        href="/"
        class="inline-flex items-center gap-2 bg-white/10 hover:bg-white/20 backdrop-blur-lg text-white px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105"
      >
        <span>←</span>
        <span>Back to Home</span>
      </a>
    </div>
  </div>
</div>

<style>
  button:disabled {
    cursor: not-allowed;
  }
</style>
