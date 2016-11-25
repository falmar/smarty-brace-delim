<script type="text/javascript">
funcion () {ldelim}
  let some = 0
  const myObject = {ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim}

{rdelim}

call({ldelim}
  hello: "world"
{rdelim}, {ldelim}
  world: "hello"
{rdelim})

let array = [{ldelim}
  hello: "world",
  myObject:{ldelim}
    one: 1,
    two: [2, 2]
  {rdelim}
{rdelim}]

inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim})
</script>
