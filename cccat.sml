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
    val fileName = case CommandLine.arguments () of
        [arg] => arg
        | _ => raise NoArgs
    val stream = TextIO.openIn fileName
in
    app(stream)
end
val _ = main ();
