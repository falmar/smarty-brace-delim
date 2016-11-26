<body>
  {$some_variable}

  Outside the script tag may be pure html or may not

<script type="text/javascript">
let myVar = {json_decode($jsonVariable)}
let myOtherVar = '{$wuuuu}'
console.log({include file=$myCustomFile})

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

const strangeObject = {ldelim}maybe: {ldelim}it: {ldelim}wont: {ldelim}work: "?"
{rdelim}, maybe: ""{rdelim}, dit: "not"{rdelim}, work: "entirely"{rdelim}

inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim})
</script>
</body>
