<script lang="ts">
  interface Props {
    children: any;
  }

  let { children }: Props = $props();
</script>

<div>
  {@render children()}
</div>
