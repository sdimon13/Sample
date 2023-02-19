<template>
  <div class="overlay">
    <form>
      <div class="con">
        <header class="head-form">
          <h2>Log In</h2>
          <p>login here using your username and password</p>
        </header>
        <br>
        <div class="field-set">
          <span class="input-item">
            <i class="fa fa-user-circle"></i>
          </span>
          <input class="form-input" id="txt-input" type="text" placeholder="@UserName" required>
          <br>
          <span class="input-item">
            <i class="fa fa-key"></i>
          </span>
          <input class="form-input" type="password" placeholder="Password" id="pwd" name="password" required>
          <span>
            <i class="fa fa-eye" aria-hidden="true" type="button" id="eye"></i>
          </span>
          <br>
          <button class="btn log-in"> Log In</button>
        </div>
        <div class="other">
          <button class="btn submits frgt-pass">Forgot Password</button>
          <button class="btn submits sign-up">Sign Up
            <i class="fa fa-user-plus" aria-hidden="true"></i>
          </button>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'LoginUser',
  data() {
    return {
      phoneNumber: '',
      password: '',
    }
  },
  methods: {
    async login() {
      try {
        const response = await axios.post('/api/login', {
          phoneNumber: this.phoneNumber,
          password: this.password,
        })
        // Handle successful login
        console.log(response.data)
      } catch (error) {
        // Handle error
        console.error(error)
      }
    },
  },
}
</script>