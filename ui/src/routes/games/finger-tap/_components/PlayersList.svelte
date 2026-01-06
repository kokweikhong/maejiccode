<script lang="ts">
  import type { Player } from '$lib/types/player';

  interface Props {
    players: Player[];
    removePlayer: (index: number) => void;
  }

  let { players, removePlayer }: Props = $props();
</script>

<div class="space-y-2 max-h-96 overflow-y-auto custom-scrollbar">
  {#if players.length === 0}
    <div class="text-center py-12 text-white/60">
      <div class="text-6xl mb-4">👥</div>
      <p>No players added yet</p>
    </div>
  {:else}
    {#each players as player, i}
      <div
        class="flex items-center justify-between bg-white/20 rounded-lg p-4 hover:bg-white/30 transition-colors"
      >
        <div class="flex items-center gap-3">
          <div
            class="w-8 h-8 rounded-full bg-linear-to-br from-yellow-400 to-orange-500 flex items-center justify-center text-white font-bold"
          >
            {i + 1}
          </div>
          <span class="text-white font-medium text-lg">{player.name}</span>
        </div>
        <button
          onclick={() => removePlayer(i)}
          class="text-red-400 hover:text-red-300 font-bold text-2xl"
          title="Remove"
        >
          ×
        </button>
      </div>
    {/each}
  {/if}
</div>

<style>
  .custom-scrollbar::-webkit-scrollbar {
    width: 8px;
  }

  .custom-scrollbar::-webkit-scrollbar-track {
    background: rgba(255, 255, 255, 0.1);
    border-radius: 10px;
  }

  .custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.3);
    border-radius: 10px;
  }

  .custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.5);
  }
</style>
