<body>
  {$some_variable}

  Outside the script tag may be pure html or may not

<script type="text/javascript">
let myVar = {json_decode($jsonVariable)}
let myOtherVar = '{$wuuuu}'
console.log({include file=$myCustomFile})
const single = {ldelim}{rdelim}

// leave this {ldelim} and {rdelim} intact
console.log('{rdelim}')
console.log("{ldelim}")
object.call('{rdelim}', "{ldelim}", `{ldelim} & {rdelim}`)
object = {ldelim}left: ["{lrdelim}", "{rdelim}"], right: {ldelim}"{rdelim}", "{ldelim}"{rdelim}{rdelim}

// this is not actually a {literal}
funcion () {ldelim}// this have ldelim: {ldelim} ?
  let some = 0
  const myObject = {ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim}

{rdelim}
// of course not the end of {/literal}

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
  {rdelim} // this must be rdelim: {rdelim}
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

function () {ldelim}/**
Everything inside
multiline comment must not be parsed!
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

const strangeObject = {ldelim}maybe: {ldelim}it: {ldelim}wont: {ldelim}work: "?"
{rdelim}, maybe: ""{rdelim}, did: "not"{rdelim}, work: "entirely"{rdelim}
*/{rdelim}

({ldelim}[{ldelim}{*
const strangeObject = {maybe: {it: {wont: {work: "?"
}, maybe: ""}, did: "not"}, work: "entirely"}
call({ldelim}
  hello: "world"
{rdelim}, {ldelim}
  world: "hello"
{rdelim})
*}{rdelim}]{rdelim})

// regexp none should be touched {$extra_regexp_pattern}
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
        return exec(/{$extra_regexp_pattern}/, value) // untouched
    default:
        return false
{rdelim}

// this {object has { lots and lots for brackets {
const strangeObject = {ldelim}maybe: {ldelim}it: {ldelim}wont: {ldelim}work: "?"
{rdelim}, maybe: ""{rdelim}, did: "not"{rdelim}, work: "entirely"{rdelim}
// but } it should not} be affected at all }

inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim})
</script>
</body>
