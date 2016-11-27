<body>
  {$some_variable}

  Outside the script tag may be pure html or may not

<script type="text/javascript">
let myVar = {json_decode($jsonVariable)}
let myOtherVar = '{$wuuuu}'
console.log({include file=$myCustomFile})
const single = {ldelim}{rdelim}

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

{literal}

$.fn.serializeObject = function () {
  var o = {}
  var a = this.serializeArray()
  $.each(a, function () {
    if (o[this.name] !== undefined) {
      if (!o[this.name].push) {
        o[this.name] = [o[this.name]]
      }
      o[this.name].push(this.value || '')
    } else {
      o[this.name] = this.value || ''
    }
  })

  return o
}

{/literal}

// regexp none should be touched
switch (key) {ldelim}
    case '_':
        return exec(/^[0-9]{11}$/, value)
    case '_':
        return exec(/^[0-9]{2}$/, value)
    case '_':
        return exec(/^[a-zA-Z]{1,2}[0-9]{2,3}$/, value)
    case '_':
        return exec(/^[0-9]{7,10}$/, value)
    case '_':
        return exec(/^\w{6}$/, value)
    default:
        return false
{rdelim}

const strangeObject = {ldelim}maybe: {ldelim}it: {ldelim}wont: {ldelim}work: "?"
{rdelim}, maybe: ""{rdelim}, did: "not"{rdelim}, work: "entirely"{rdelim}

inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim})
</script>
</body>
