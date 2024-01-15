<script>
  import Numbers from '$lib/components/Numbers.svelte';

  export let data;
</script>

<Numbers numbers=7 max=35 pick={ data.numbers }/>