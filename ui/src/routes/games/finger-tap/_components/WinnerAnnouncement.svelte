<script lang="ts">
  import type { Player } from '$lib/types/player';

  interface Props {
    winners: Player[];
    maxTaps: number;
  }

  let { winners, maxTaps }: Props = $props();
</script>

<div
  class="bg-white/10 backdrop-blur-lg rounded-2xl p-8 text-center animate-scaleIn"
>
  <div class="text-6xl mb-4">🏆</div>
  <h2 class="text-4xl font-bold text-white mb-4">
    {#if winners.length === 1}
      Winner: {winners[0].name}!
    {:else if winners.length > 1}
      It's a Tie!
    {:else}
      No Taps Recorded
    {/if}
  </h2>
  {#if winners.length > 0}
    <p class="text-2xl text-yellow-400 font-bold">
      {maxTaps} tap{maxTaps !== 1 ? 's' : ''}
    </p>
  {/if}
</div>

<style>
  .animate-scaleIn {
    animation: scaleIn 0.5s ease-out;
  }
</style>
