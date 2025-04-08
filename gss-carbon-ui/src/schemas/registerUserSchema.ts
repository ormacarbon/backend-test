import * as yup from "yup";

export const registerUserSchema = yup
  .object({
    name: yup.string().required("Name is required"),
    email: yup.string().email("Email is invalid").required("Email is required"),
    phone: yup
      .string()
      .required("Phone number is required")
      .matches(
        /^\d{11}$/,
        "Phone number format is invalid. Ex: 47999999999",
      ),
  })
  .required();

export type SignUpSchemaType = yup.InferType<typeof registerUserSchema>;
