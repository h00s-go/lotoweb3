<script>
  import Lottery from '$lib/components/Lottery.svelte';

  export let data;
</script>

<Lottery numbers=7 max=35 pick={ data.numbers }/>