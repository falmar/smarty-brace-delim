<body>
  {$some_variable}

  Outside the script tag may be pure html or may not

<script type="text/javascript">
let myVar = {json_decode($jsonVariable)}
let myOtherVar = '{$wuuuu}'
console.log({include file=$myCustomFile})
const single = {}

funcion () {
  let some = 0
  const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]}}

}

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
  }
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

const strangeObject = {maybe: {it: {wont: {work: "?"
}, maybe: ""}, did: "not"}, work: "entirely"}

inline_call({hello: "world", myObject:{one: 1, two: [2, 2]}})
</script>
</body>
