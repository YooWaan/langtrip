import java.lang.annotation.ElementType;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import java.lang.annotation.Retention;
import java.lang.reflect.Method;

// declare
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD) // method only
public @interface Option {
    boolean isActive() default false;
}

class Api {
  @Option(isActive = true)
  public String get() {
     return "get";
  }
  @Option
  public String put() {
     return "put";
  }
}

class Caller {
    public void call(Object instance, String methodName) {
        try {
            Method method = instance.getClass().getMethod(methodName);
            Option option = method.getAnnotation(Option.class);
            String result = option.isActive()
                ? String.format("warp[%s]", method.invoke(instance))
                : method.invoke(instance).toString();
            
            System.out.println("Result[" + result + "]");
        } catch (Throwable cause) {
            cause.printStackTrace();
            throw new RuntimeException(cause);
        }
    }
}

Api api = new Api();
new Caller().call(api, "get");
new Caller().call(api, "put");

/exit
