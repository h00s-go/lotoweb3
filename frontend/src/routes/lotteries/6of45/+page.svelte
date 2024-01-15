<script>
  import Lottery from '$lib/components/Lottery.svelte';

  export let data;
</script>

<Lottery lottery={ data.lottery } />
