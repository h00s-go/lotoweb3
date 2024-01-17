<script>
  export let lottery;
</script>

<div class="card">
  <div class="card-header">
    <b>{ lottery.numbers } / { lottery.max }</b>
  </div>
  <div class="card-body">
    <table class="table table-hover">
      <thead>
        <tr>
          { #each lottery.picks[0] as number, index }
            <th>#{ index + 1 }</th>
          { /each }
        </tr>
      </thead>
      <tbody>
        { #each lottery.picks as pick }
          <tr>
            { #each pick as number }
              <td>{ number }</td>
            { /each }
          </tr>
        { /each }
      </tbody>
    </table>
  </div>
</div>