export default function ({ $axios, redirect }) {
  console.log($axios)
  $axios.setToken('access_token')

  $axios.onResponse((config) => {
    $axios.setHeader('Access-Control-Allow-Origin', '*')
  })
}
