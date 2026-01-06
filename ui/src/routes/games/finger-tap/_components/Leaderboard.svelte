<script lang="ts">
  import type { Player } from '$lib/types/player';

  interface Props {
    sortedPlayers: Player[];
  }
  let { sortedPlayers }: Props = $props();
</script>

<div class="bg-white/10 backdrop-blur-lg rounded-2xl p-6">
  <h3 class="text-2xl font-bold text-white mb-4 text-center">📊 Results</h3>
  <div class="space-y-3">
    {#each sortedPlayers as player, i}
      <div
        class="flex items-center justify-between bg-white/20 rounded-lg p-4 {i ===
          0 && player.taps > 0
          ? 'ring-4 ring-yellow-400'
          : ''}"
      >
        <div class="flex items-center gap-3">
          <div class="text-3xl">
            {#if i === 0 && player.taps > 0}
              🥇
            {:else if i === 1 && player.taps > 0}
              🥈
            {:else if i === 2 && player.taps > 0}
              🥉
            {:else}
              {i + 1}
            {/if}
          </div>
          <span class="text-white font-bold text-xl">{player.name}</span>
        </div>
        <div class="text-yellow-400 font-bold text-2xl">
          {player.taps}
        </div>
      </div>
    {/each}
  </div>
</div>
