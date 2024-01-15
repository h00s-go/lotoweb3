<script>
  export let lottery;
</script>

<div class="card">
  <div class="card-header">
    <b>{ lottery.numbers } / { lottery.max }</b>
  </div>
  <div class="card-body">
    { #each lottery.picks as pick }
      <p class="pb-2">{ pick.join(', ') }</p>
    { /each }
  </div>
</div>