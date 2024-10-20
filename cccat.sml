exception NoArgs

fun app(stream) =
    case (TextIO.endOfStream(stream), stream) of
        (true, _) => print("")
        | (false, stream) => 
            (
                (case TextIO.input1(stream) of
                    SOME c => print(String.str(c))
                    | NONE => print("\n"));
             app(stream)
             )
fun main () =

let
    val arguments = case CommandLine.arguments () of
        [] => raise NoArgs
        | args => args
    val streams = List.map (fn (x) => if x = "-" then TextIO.stdIn else TextIO.openIn x) arguments
in
    List.app app streams
end
val _ = main ();
