<script>
  import Lottery from '$lib/components/Lottery.svelte';

  const { data } = $props();
</script>

<Lottery lottery={ data.lottery } />
