let container = document.querySelector('.container');
let login = document.querySelector('.login-btn');
let register = document.querySelector('.register-btn');
register.addEventListener('click', () => {
    container.classList.add('active');
})
login.addEventListener('click', () => {
    container.classList.remove('active');
})


let formlogin = document.getElementById("login-form")
let formregister = document.getElementById("register-form")
const newHeaders = new Headers();
newHeaders.append("Content-Type", "application/json"); 



formlogin.addEventListener('submit', (e) => {
    e.preventDefault();

    const formdata = {
        form: formlogin,
        url: "/loginInfo",
        redirect: "/", 
    };

    handleFormSubmit(formdata);
});


formregister.addEventListener('submit', (e) => {
    e.preventDefault();

    const formdata = {
        form: formregister,
        url: "/registerInfo",
        redirect: "/login", 
    };

    handleFormSubmit(formdata);
});


function handleFormSubmit(formdata) {
    const data = new FormData(formdata.form);

    fetch(formdata.url, {
        method: "POST",
        body: data,
    })
        .then((response) => response.json())
        .then((result) => {
            console.log(result);

            if (result.success) {
       
                window.location.href = formdata.redirect;
            } else {
            
                displayErrorMessage(formdata.form, result.message+" please try again");
            }
        })
        .catch((error) => {
            console.error("Error:", error);
        });
}

function displayErrorMessage(form, message) {
   
    let existingError = form.querySelector(".error-message");
    if (existingError) {
        existingError.textContent = message;
    } else {
      
        let perror = document.createElement("p");
        perror.textContent = message;
        perror.style.color = "red";
        perror.style.fontSize = "15px";
        perror.style.fontWeight = "bold";
        perror.style.textAlign = "center";
        perror.classList.add("error-message"); 
        let existingError = form.querySelector(".error-message");
        if (existingError) {
            existingError.remove();
            form.appendChild(perror);
        }else{
            form.appendChild(perror);
        }
    }
}