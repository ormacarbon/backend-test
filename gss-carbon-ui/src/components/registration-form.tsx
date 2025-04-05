import { useRegisterUser } from "@/hooks/useUser";
import { yupResolver } from "@hookform/resolvers/yup";
import { Loader2 } from "lucide-react";
import React from "react";
import { useForm } from "react-hook-form";
import { Button } from "../components/ui/button";
import { Input } from "../components/ui/input";
import { Label } from "../components/ui/label";
import {
  SignUpSchemaType,
  registerUserSchema,
} from "../schemas/registerUserSchema";

interface RegistrationFormProps {
  referralCode: string | null;
}

const RegistrationForm: React.FC<RegistrationFormProps> = ({
  referralCode,
}) => {
  const { mutateAsync, isPending } = useRegisterUser();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignUpSchemaType>({
    resolver: yupResolver(registerUserSchema),
  });

  const onSubmit = async (data: SignUpSchemaType) => {
    await mutateAsync({ ...data, referralCode });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
      <div className="space-y-2">
        <Label htmlFor="name" className="text-slate-700">
          Full Name
        </Label>
        <Input
          id="name"
          placeholder="Enter your full name"
          {...register("name")}
          className={`border-slate-200 focus:border-emerald-500 focus:ring-emerald-500 ${errors.name ? "border-red-300 ring-1 ring-red-300" : ""}`}
          disabled={isPending}
        />
        {errors.name && (
          <p className="mt-1 text-xs text-red-500">{errors.name.message}</p>
        )}
      </div>

      <div className="space-y-2">
        <Label htmlFor="email" className="text-slate-700">
          Email Address
        </Label>
        <Input
          id="email"
          type="email"
          placeholder="Enter your email address"
          {...register("email")}
          className={`border-slate-200 focus:border-emerald-500 focus:ring-emerald-500 ${
            errors.email ? "border-red-300 ring-1 ring-red-300" : ""
          }`}
          disabled={isPending}
        />
        {errors.email && (
          <p className="mt-1 text-xs text-red-500">{errors.email.message}</p>
        )}
      </div>

      <div className="space-y-2">
        <Label htmlFor="phone" className="text-slate-700">
          Phone Number
        </Label>
        <Input
          id="phone"
          type="tel"
          placeholder="Enter your phone number"
          {...register("phone")}
          className={`border-slate-200 focus:border-emerald-500 focus:ring-emerald-500 ${
            errors.phone ? "border-red-300 ring-1 ring-red-300" : ""
          }`}
          disabled={isPending}
        />
        {errors.phone && (
          <p className="mt-1 text-xs text-red-500">{errors.phone.message}</p>
        )}
      </div>

      <Button
        type="submit"
        className="w-full bg-gradient-to-r from-emerald-500 to-teal-500 text-white hover:from-emerald-600 hover:to-teal-600"
        disabled={isPending}
      >
        {isPending ? (
          <>
            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            Registering...
          </>
        ) : (
          "Register Now"
        )}
      </Button>
    </form>
  );
};

export default RegistrationForm;
