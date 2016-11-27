<body>
  {$some_variable}

  Outside the script tag may be pure html or may not

<script type="text/javascript">
let myVar = {json_decode($jsonVariable)}
let myOtherVar = '{$wuuuu}'
console.log({include file=$myCustomFile})
const single = {}

// this is not actually a {literal}
funcion () {// this have ldelim: {ldelim} ?
  let some = 0
  const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]}}

}
// of course not the end of {/literal}

call({
  hello: "world"
}, {
  world: "hello"
})

let array = [{
  hello: "world",
  myObject:{
    one: 1,
    two: [2, 2]
  } // this must be rdelim: {rdelim}
}]

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

function () {/**
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
*/}

({[{{*
const strangeObject = {maybe: {it: {wont: {work: "?"
}, maybe: ""}, did: "not"}, work: "entirely"}
call({ldelim}
  hello: "world"
{rdelim}, {ldelim}
  world: "hello"
{rdelim})
*}}]})

// regexp none should be touched {$extra_regexp_pattern}
switch (key) {
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
}

// this {object has { lots and lots for brackets {
const strangeObject = {maybe: {it: {wont: {work: "?"
}, maybe: ""}, did: "not"}, work: "entirely"}
// but } it should not} be affected at all }

inline_call({hello: "world", myObject:{one: 1, two: [2, 2]}})
</script>
</body>
