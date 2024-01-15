<script>
  import Numbers from '$lib/components/Numbers.svelte';

  export let data;
</script>

<Numbers numbers=6 max=45 pick={ data.numbers }/>
