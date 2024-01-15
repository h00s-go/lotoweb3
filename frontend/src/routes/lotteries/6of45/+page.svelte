<script>
  import Lottery from '$lib/components/Lottery.svelte';

  export let data;
</script>

<Lottery numbers=6 max=45 pick={ data.numbers }/>
