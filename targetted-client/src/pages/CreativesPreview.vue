<template>
      <div>
      <div id="content-wrapper">
        <div id="content">
          <b-row class="app-new-steps-header">
              <b-col>
                <h2>Реклама в сториз</h2>

                <br><br><br>...<br><br><br><br>

                <h2>Реклама в ленте</h2>

                <div class="hidden">
                <svg id="dots" viewBox="0 0 48 48">
                    <circle clip-rule="evenodd" cx="8" cy="24" fill-rule="evenodd" r="4.5"></circle>
                    <circle clip-rule="evenodd" cx="24" cy="24" fill-rule="evenodd" r="4.5"></circle>
                    <circle clip-rule="evenodd" cx="40" cy="24" fill-rule="evenodd" r="4.5"></circle>
                </svg>

                <svg id="like" viewBox="0 0 48 48">
                    <path d="M34.6 6.1c5.7 0 10.4 5.2 10.4 11.5 0 6.8-5.9 11-11.5 16S25 41.3 24 41.9c-1.1-.7-4.7-4-9.5-8.3-5.7-5-11.5-9.2-11.5-16C3 11.3 7.7 6.1 13.4 6.1c4.2 0 6.5 2 8.1 4.3 1.9 2.6 2.2 3.9 2.5 3.9.3 0 .6-1.3 2.5-3.9 1.6-2.3 3.9-4.3 8.1-4.3m0-3c-4.5 0-7.9 1.8-10.6 5.6-2.7-3.7-6.1-5.5-10.6-5.5C6 3.1 0 9.6 0 17.6c0 7.3 5.4 12 10.6 16.5.6.5 1.3 1.1 1.9 1.7l2.3 2c4.4 3.9 6.6 5.9 7.6 6.5.5.3 1.1.5 1.6.5.6 0 1.1-.2 1.6-.5 1-.6 2.8-2.2 7.8-6.8l2-1.8c.7-.6 1.3-1.2 2-1.7C42.7 29.6 48 25 48 17.6c0-8-6-14.5-13.4-14.5z"></path>
                </svg>

                <svg id="dislike" viewBox="0 0 48 48">
                    <path d="M34.6 3.1c-4.5 0-7.9 1.8-10.6 5.6-2.7-3.7-6.1-5.5-10.6-5.5C6 3.1 0 9.6 0 17.6c0 7.3 5.4 12 10.6 16.5.6.5 1.3 1.1 1.9 1.7l2.3 2c4.4 3.9 6.6 5.9 7.6 6.5.5.3 1.1.5 1.6.5s1.1-.2 1.6-.5c1-.6 2.8-2.2 7.8-6.8l2-1.8c.7-.6 1.3-1.2 2-1.7C42.7 29.6 48 25 48 17.6c0-8-6-14.5-13.4-14.5z"></path>
                </svg>
                <svg id="comment" viewBox="0 0 48 48">
                    <path clip-rule="evenodd" d="M47.5 46.1l-2.8-11c1.8-3.3 2.8-7.1 2.8-11.1C47.5 11 37 .5 24 .5S.5 11 .5 24 11 47.5 24 47.5c4 0 7.8-1 11.1-2.8l11 2.8c.8.2 1.6-.6 1.4-1.4zm-3-22.1c0 4-1 7-2.6 10-.2.4-.3.9-.2 1.4l2.1 8.4-8.3-2.1c-.5-.1-1-.1-1.4.2-1.8 1-5.2 2.6-10 2.6-11.4 0-20.6-9.2-20.6-20.5S12.7 3.5 24 3.5 44.5 12.7 44.5 24z" fill-rule="evenodd"></path>
                </svg>

                <svg id="send" viewBox="0 0 48 48">
                    <path d="M47.8 3.8c-.3-.5-.8-.8-1.3-.8h-45C.9 3.1.3 3.5.1 4S0 5.2.4 5.7l15.9 15.6 5.5 22.6c.1.6.6 1 1.2 1.1h.2c.5 0 1-.3 1.3-.7l23.2-39c.4-.4.4-1 .1-1.5zM5.2 6.1h35.5L18 18.7 5.2 6.1zm18.7 33.6l-4.4-18.4L42.4 8.6 23.9 39.7z"></path>
                </svg>

                <svg id="save" viewBox="0 0 48 48">
                    <path d="M43.5 48c-.4 0-.8-.2-1.1-.4L24 29 5.6 47.6c-.4.4-1.1.6-1.6.3-.6-.2-1-.8-1-1.4v-45C3 .7 3.7 0 4.5 0h39c.8 0 1.5.7 1.5 1.5v45c0 .6-.4 1.2-.9 1.4-.2.1-.4.1-.6.1zM24 26c.8 0 1.6.3 2.2.9l15.8 16V3H6v39.9l15.8-16c.6-.6 1.4-.9 2.2-.9z"></path>
                </svg>

                <svg id="smiley" viewBox="0 0 48 48">
                    <path d="M24 48C10.8 48 0 37.2 0 24S10.8 0 24 0s24 10.8 24 24-10.8 24-24 24zm0-45C12.4 3 3 12.4 3 24s9.4 21 21 21 21-9.4 21-21S35.6 3 24 3z"></path>
                    <path d="M34.9 24c0-1.4-1.1-2.5-2.5-2.5s-2.5 1.1-2.5 2.5 1.1 2.5 2.5 2.5 2.5-1.1 2.5-2.5zm-21.8 0c0-1.4 1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5-1.1 2.5-2.5 2.5-2.5-1.1-2.5-2.5zM24 37.3c-5.2 0-8-3.5-8.2-3.7-.5-.6-.4-1.6.2-2.1.6-.5 1.6-.4 2.1.2.1.1 2.1 2.5 5.8 2.5 3.7 0 5.8-2.5 5.8-2.5.5-.6 1.5-.7 2.1-.2.6.5.7 1.5.2 2.1 0 .2-2.8 3.7-8 3.7z"></path>
                </svg>
            </div>

        <article class="instapost">
            <header class="instapost__header">
                <a class="profile-img" href="#" target="_blank">
                    <img src="@src/assets/Ellipse64.png" alt="username" />
                </a>
                <div class="profile-name">
                    <a class="user instalink" href="" target="_blank">
                        username
                    </a>
                    <a class="location" href="#" target="_blank">Debrecen, New York</a>
                </div>
                <button class="btn btn-more">
                    <svg>
                        <use xlink:href="#dots" />
                    </svg>
                </button>
            </header>
            <section class="instapost__image">
                <img class="img img-0 show" src="@src/assets/image2.png" alt="image" />
                <div class="like-heart">
                    <svg>
                        <use xlink:href="#dislike" />
                    </svg>
                </div>
            </section>
            <section class="instapost__action">
                <button class="btn btn-like">
                    <svg class="like-icon">
                        <use xlink:href="#like" />
                    </svg>
                    <svg class="dislike-icon">
                        <use xlink:href="#dislike" />
                    </svg>
                </button>
                <button class="btn btn-comment">
                    <svg>
                        <use xlink:href="#comment" />
                    </svg>
                </button>
                <button class="btn btn-send">
                    <svg>
                        <use xlink:href="#send" />
                    </svg>
                </button>
                <button class="btn btn-save">
                    <svg>
                        <use xlink:href="#save" />
                    </svg>
                </button>
            </section>
            <section class="instapost__likes">
                Liked by <a class="instalink" href="#" target="_blank">targetted</a> and <a class="instalink" href="" target="_blank">others</a>
            </section>
            <section class="instapost__description">
                <a class="user instalink" href="#" target="_blank">
                        username
                    </a> Taking an image, freezing a moment, reveals how rich reality truly is
            </section>
            <a class="instapost__comment-list" href="#">
                View all 817 comments
            </a>
            <section class="instapost__timestamp">
                10 minutes ago
            </section>
            <!-- <section class="instapost__add-comment">
                <button class="btn btn-smiley">
                    <svg>
                        <use xlink:href="#smiley" />
                    </svg>
                </button>
                <textarea class="comment-input" placeholder="Add a comment..." rows="1"></textarea>
                <button class="btn btn-send-comment" disabled="disabled">
                    Send
                </button>
            </section> -->
        </article>

            <b-button type="button"
                    class="app-new-submit-button mt-5"
                    @click="$router.back();">
                  Назад
            </b-button>
            </b-col>
          </b-row>
        </div>
    </div>
  </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@400;700&display=swap");

* {
  box-sizing: border-box;
}

.hidden {
  width: 0;
  height: 0;
  visibility: hidden;
  display: none;
  overflow: hidden;
}

html {
  font-size: 16px;
}

article {
  display: flex;
  flex-direction: column;
  font-family: "Poppins", sans-serif;
  position: relative;
  background-color: #fbfcfc;
}

.challenge-title {
  text-align: center;
  padding: 0;
  margin: 1rem 0 0;
  color: #2f363d;
  font-size: 2rem;
  font-weight: bold;
}

.challenge-subtitle {
  text-align: center;
  margin: 0;
  font-size: 1.125rem;
  font-weight: 400;
  color: #b5bbc3;
}

.challenge-part-of {
  text-align: center;
  margin: 1rem 0 0;
  font-size: 0.875rem;
  color: #b5bbc3;
  position: fixed;
  bottom: 1rem;
  right: 1rem;
}

.challenge-part-of a,
.challenge-part-of a:visited {
  color: #b5bbc3;
  -webkit-text-decoration-style: wavy;
          text-decoration-style: wavy;
}

.challenge-part-of a:hover, .challenge-part-of a:active,
.challenge-part-of a:visited:hover,
.challenge-part-of a:visited:active {
  color: #d1d5da;
}

article .btn {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: transparent;
  border-color: transparent;
  cursor: pointer;
  width: 2.5rem;
  height: 2.5rem;
  padding: 0.5rem;
}

article .btn:focus {
  outline: 1px dotted #98a1ad;
}
article .btn svg {
  fill: #8a94a1;
  width: 100%;
  height: 100%;
}

.instalink,
.instalink:visited {
  margin-bottom: 0.25rem;
  color: #2f363d;
  text-decoration: none;
  font-weight: bold;
}
.instalink:hover, .instalink:active,
.instalink:visited:hover,
.instalink:visited:active {
  text-decoration: underline;
}

.instapost {
  margin: 1rem;
  width: 600px;
  background-color: white;
  display: flex;
  flex-direction: column;
  line-height: 1.125rem;
}

@media (max-width: 600px) {
    .instapost {
        margin: 0;
        width: 90%;
    }
}

.instapost__header {
  display: flex;
  align-items: center;
  padding: 1rem;
}

.instapost__header .profile-img img {
  -o-object-fit: cover;
     object-fit: cover;
  width: 32px;
  height: 32px;
  border-radius: 50%;
}

.instapost__header .profile-name {
  display: flex;
  flex-direction: column;
  margin-left: 0.875rem;
}

.instapost__header .profile-name a {
  font-size: 0.875rem;
}

.instapost__header .profile-name a.location,
.instapost__header .profile-name a.location:active {
  color: #5b6977;
  text-decoration: none;
}

.instapost__header .btn-more {
  margin: 0 0 0 auto;
}

.instapost__image {
  position: relative;
}

.instapost__image .img {
  -o-object-fit: cover;
     object-fit: cover;
  width: 100%;
  display: none;
  border-radius: 6px;
}

.instapost__image .img.show {
  display: block;
}

.instapost__image .like-heart {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: none;
}

.instapost__image .like-heart svg {
  fill: white;
  filter: drop-shadow(0 0 0.5rem #2f363d);
}

.instapost__image .like-heart.show {
  display: block;
  -webkit-animation: heartbeatAlt 0.8s linear forwards;
          animation: heartbeatAlt 0.8s linear forwards;
}

.instapost__action {
  display: flex;
  align-items: center;
  padding: 0 1rem;
}

.instapost__action .btn svg {
  fill: #2f363d;
}

.instapost__action .btn-like .like-icon {
  display: block;
}

.instapost__action .btn-like .dislike-icon {
  display: none;
  fill: #ed4956;
}

.instapost__action .btn-like.dislike .like-icon {
  display: none;
}

.instapost__action .btn-save {
  margin-left: auto;
}

.instapost__likes {
  color: #5b6977;
  font-size: 0.875rem;
  padding: 0 1rem;
  margin: 0 0 0.5rem;
}

.instapost__description {
  color: #5b6977;
  font-size: 0.875rem;
  padding: 0 1rem;
}

.instapost__comment-list, .instapost__comment-list:visited {
  color: #728393;
  font-size: 0.875rem;
  padding: 0.5rem 1rem;
  text-decoration: none;
}

.instapost__timestamp {
  color: #728393;
  text-transform: uppercase;
  font-size: 0.625rem;
  padding: 0 1rem 0.25rem;
}

.instapost__add-comment {
  display: flex;
  align-items: center;
  padding: 0.25rem 1rem 0;
  border-top: 1px solid #edeff1;
}

.instapost__add-comment .btn-smiley svg {
  fill: #2f363d;
}

.instapost__add-comment textarea {
  flex: 1;
  border: 0;
  resize: none;
}

.instapost__add-comment textarea:focus {
  outline: 0;
}

.instapost__add-comment .btn-send-comment {
  color: #0095f6;
  width: auto;
  font-size: 0.875rem;
  font-weight: bold;
}

.instapost__add-comment .btn-send-comment[disabled=disabled] {
  cursor: default;
  opacity: 0.3;
}
</style>
