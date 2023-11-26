function showToast(message, duration = 3000) {
  const container = document.getElementById("toast-container");
  const toast = document.createElement("div");
  toast.classList.add(
    "bg-gray-800",
    "border-t-4",
    "border-red-400",
    "text-white",
    "p-3",
    "rounded",
    "shadow",
    "opacity-0",
    "transform",
    "translate-y-8",
    "transition",
    "duration-300",
    "ease-out"
  );
  toast.textContent = message;

  container.appendChild(toast);

  setTimeout(() => {
    toast.classList.remove("opacity-0", "translate-y-8");
    toast.classList.add("opacity-100", "translate-y-0");
  }, 100);

  setTimeout(() => {
    toast.classList.remove("opacity-100", "translate-y-0");
    toast.classList.add("opacity-0", "translate-y-8");
    toast.addEventListener("transitionend", () => toast.remove());
  }, duration);
}

function getLoadingSpinner() {
  return document.getElementById("loading-spinner");
}

function hideLoadingSpinner(spinner) {
  spinner.classList.add("hidden");
}

function createKeydownHandler(submitButton, emailInput) {
  return function (event) {
    resetToInitial(submitButton, emailInput);
    emailInput.removeEventListener("keydown", emailInput.keydownHandler);
  };
}

function resetToInitial(submitButton, emailInput) {
  submitButton.classList.remove("bg-green-500", "bg-[#f1aa27]");
  submitButton.classList.add("bg-[#276ef1]");
  submitButton.value = "Subscribe";
  emailInput.classList.remove("hidden");
}

function setToGreen(submitButton) {
  localStorage.setItem("subscriptionStatus", true);
  submitButton.classList.remove("bg-[#276ef1]");
  submitButton.classList.remove("hidden");
  submitButton.classList.add("bg-green-500");
  submitButton.value = "✅ Subscribed";
}

function showInput(inputButton, emailInput) {
  inputButton.classList.remove("hidden");
  emailInput.classList.remove("hidden");
}

function hideInput(inputButton, emailForm) {
  inputButton.classList.add("hidden");
  emailForm.classList.remove("space-x-4");
}

function setToRed(submitButton) {
  submitButton.classList.remove("bg-[#276ef1]");
  submitButton.classList.add("bg-[#f1aa27]");
  submitButton.value = "❌ Failed";
}

function getEmailForm() {
  return document.getElementById("email-form");
}

function getEmailInput() {
  return document.getElementById("email-input");
}

function getSubmitButton() {
  return document.getElementById("submit-button");
}

function checkRegistration() {
  const subscriptionStatus = localStorage.getItem("subscriptionStatus");
  const emailInput = getEmailInput();
  const loadingSpinner = getLoadingSpinner();
  const submitButton = getSubmitButton();

  if (subscriptionStatus) {
    const inputForm = getEmailForm();

    setToGreen(submitButton);
    hideLoadingSpinner(loadingSpinner);
    hideInput(emailInput, inputForm);
    return;
  }
  showInput(emailInput, submitButton);
  hideLoadingSpinner(loadingSpinner);
}

function submitForm(event) {
  event.preventDefault();

  const emailInput = getEmailInput();
  const email = emailInput.value;
  const inputForm = getEmailForm();

  const jsonData = JSON.stringify({ email });

  const submitButton = getSubmitButton();

  fetch("/api/v1/subscribe", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: jsonData,
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.status === "success") {
        setToGreen(submitButton);
        hideInput(emailInput, inputForm);
      } else if (data.message && data.message === "Email already exists") {
        setToRed(submitButton);
        emailInput.keydownHandler = createKeydownHandler(
          submitButton,
          emailInput
        );
        emailInput.addEventListener("keydown", emailInput.keydownHandler);
        showToast("This email is already registered.");
      } else {
        setToRed(submitButton);
        emailInput.keydownHandler = createKeydownHandler(
          submitButton,
          emailInput
        );
        emailInput.addEventListener("keydown", emailInput.keydownHandler);
        showToast("Failed to subscribe. Please try again.");
      }
    })
    .catch((error) => {
      setToRed(submitButton);
      emailInput.keydownHandler = createKeydownHandler(
        submitButton,
        emailInput
      );
      emailInput.addEventListener("keydown", emailInput.keydownHandler);
      showToast("Failed to subscribe. Please try again.");
    });
}

addEventListener("load", (event) => {
  checkRegistration();
});
